package main

import (
	"testing"
)

type dummyConsumer struct {
}

func (consumer dummyConsumer) Consume(interface{}) {
	// Do nothing, only consume
}

func TestNewBrokerSuccessful(t *testing.T) {
	consumer := dummyConsumer{}
	NewBroker(consumer)
}
