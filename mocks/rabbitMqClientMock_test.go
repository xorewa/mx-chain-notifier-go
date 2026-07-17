package mocks

import (
	"fmt"
	"sync"
	"testing"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/require"
)

func TestRabbitClientMockEntriesAreDeeplyDetached(t *testing.T) {
	client := NewRabbitClientMock()
	original := amqp.Publishing{
		Body: []byte("payload"),
		Headers: amqp.Table{
			"bytes":  []byte("header"),
			"nested": amqp.Table{"value": []byte("nested-value")},
			"array":  []interface{}{[]byte("array-value")},
		},
	}

	require.NoError(t, client.Publish("events", "", false, false, original))

	// Mutating the caller-owned input after Publish must not alter mock state.
	original.Body[0] = 'X'
	original.Headers["bytes"].([]byte)[0] = 'X'

	first := client.GetEntries()
	firstEvent := first["events"]
	require.Equal(t, []byte("payload"), firstEvent.Body)
	require.Equal(t, []byte("header"), firstEvent.Headers["bytes"])

	// Mutating a returned snapshot, including nested reference fields, must not
	// alter a later snapshot.
	firstEvent.Body[0] = 'Y'
	firstEvent.Headers["bytes"].([]byte)[0] = 'Y'
	firstEvent.Headers["nested"].(amqp.Table)["value"].([]byte)[0] = 'Y'
	firstEvent.Headers["array"].([]interface{})[0].([]byte)[0] = 'Y'
	first["events"] = firstEvent

	second := client.GetEntries()["events"]
	require.Equal(t, []byte("payload"), second.Body)
	require.Equal(t, []byte("header"), second.Headers["bytes"])
	require.Equal(t, []byte("nested-value"), second.Headers["nested"].(amqp.Table)["value"])
	require.Equal(t, []byte("array-value"), second.Headers["array"].([]interface{})[0])
}

func TestRabbitClientMockConcurrentWritesAndSnapshots(t *testing.T) {
	client := NewRabbitClientMock()

	var waitGroup sync.WaitGroup
	for worker := 0; worker < 8; worker++ {
		worker := worker
		waitGroup.Add(2)

		go func() {
			defer waitGroup.Done()
			for index := 0; index < 200; index++ {
				_ = client.Publish(
					fmt.Sprintf("exchange-%d-%d", worker, index),
					"",
					false,
					false,
					amqp.Publishing{Body: []byte("payload")},
				)
			}
		}()

		go func() {
			defer waitGroup.Done()
			for index := 0; index < 200; index++ {
				entries := client.GetEntries()
				entries[fmt.Sprintf("snapshot-%d-%d", worker, index)] = amqp.Publishing{}
			}
		}()
	}

	waitGroup.Wait()
	require.Len(t, client.GetEntries(), 8*200)
}
