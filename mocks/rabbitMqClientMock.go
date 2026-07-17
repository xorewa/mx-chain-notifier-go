package mocks

import (
	"sync"

	"github.com/streadway/amqp"
)

// RabbitClientMock -
type RabbitClientMock struct {
	mut    sync.RWMutex
	events map[string]amqp.Publishing
}

// NewRabbitClientMock -
func NewRabbitClientMock() *RabbitClientMock {
	return &RabbitClientMock{
		events: make(map[string]amqp.Publishing),
	}
}

// Publish -
func (rc *RabbitClientMock) Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	rc.mut.Lock()
	defer rc.mut.Unlock()

	rc.events[exchange] = clonePublishing(msg)

	return nil
}

// ExchangeDeclare -
func (rc *RabbitClientMock) ExchangeDeclare(name, kind string) error {
	return nil
}

// ConnErrChan -
func (rc *RabbitClientMock) ConnErrChan() chan *amqp.Error {
	return make(chan *amqp.Error)
}

// CloseErrChan -
func (rc *RabbitClientMock) CloseErrChan() chan *amqp.Error {
	return make(chan *amqp.Error)
}

// Reconnect -
func (rc *RabbitClientMock) Reconnect() {
}

// ReopenChannel -
func (rc *RabbitClientMock) ReopenChannel() {
}

// GetEntries -
func (rc *RabbitClientMock) GetEntries() map[string]amqp.Publishing {
	rc.mut.RLock()
	defer rc.mut.RUnlock()

	entries := make(map[string]amqp.Publishing, len(rc.events))
	for exchange, event := range rc.events {
		entries[exchange] = clonePublishing(event)
	}

	return entries
}

func clonePublishing(event amqp.Publishing) amqp.Publishing {
	cloned := event
	cloned.Body = append([]byte(nil), event.Body...)
	cloned.Headers = cloneAMQPTable(event.Headers)

	return cloned
}

func cloneAMQPTable(table amqp.Table) amqp.Table {
	if table == nil {
		return nil
	}

	cloned := make(amqp.Table, len(table))
	for key, value := range table {
		cloned[key] = cloneAMQPValue(value)
	}

	return cloned
}

func cloneAMQPValue(value interface{}) interface{} {
	switch typed := value.(type) {
	case amqp.Table:
		return cloneAMQPTable(typed)
	case []byte:
		return append([]byte(nil), typed...)
	case []interface{}:
		cloned := make([]interface{}, len(typed))
		for index, entry := range typed {
			cloned[index] = cloneAMQPValue(entry)
		}
		return cloned
	default:
		return value
	}
}

// Close -
func (rc *RabbitClientMock) Close() {
}

// IsInterfaceNil -
func (rc *RabbitClientMock) IsInterfaceNil() bool {
	return rc == nil
}
