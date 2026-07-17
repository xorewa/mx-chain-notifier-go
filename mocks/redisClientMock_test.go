package mocks

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRedisClientMockEntriesAreDetached(t *testing.T) {
	client := NewRedisClientMock()
	_, err := client.SetEntry(context.Background(), "first", true, time.Minute)
	require.NoError(t, err)

	first := client.GetEntries()
	delete(first, "first")
	first["injected"] = true

	second := client.GetEntries()
	require.Equal(t, map[string]bool{"first": true}, second)
}

func TestRedisClientMockConcurrentWritesAndSnapshots(t *testing.T) {
	client := NewRedisClientMock()

	var waitGroup sync.WaitGroup
	for worker := 0; worker < 8; worker++ {
		worker := worker
		waitGroup.Add(2)

		go func() {
			defer waitGroup.Done()
			for index := 0; index < 200; index++ {
				_, _ = client.SetEntry(
					context.Background(),
					fmt.Sprintf("entry-%d-%d", worker, index),
					true,
					time.Minute,
				)
			}
		}()

		go func() {
			defer waitGroup.Done()
			for index := 0; index < 200; index++ {
				entries := client.GetEntries()
				entries[fmt.Sprintf("snapshot-%d-%d", worker, index)] = true
			}
		}()
	}

	waitGroup.Wait()
	require.Len(t, client.GetEntries(), 8*200)
}
