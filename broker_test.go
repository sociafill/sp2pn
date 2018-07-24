package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type dummyConsumer struct {
}

func (consumer dummyConsumer) Consume(Observable, interface{}) {
	// Do nothing, only consume
}

type dummyObservable struct {
	launchesCount int
}

func (observable *dummyObservable) Identifier() Identifier {
	return "dummy"
}

func (observable *dummyObservable) Poll() []interface{} {
	observable.launchesCount++
	var result []interface{}
	return result
}

func TestNewBrokerSuccessful(t *testing.T) {
	consumer := dummyConsumer{}
	NewBroker(consumer)
}

func TestPollingIsSpawnedSuccessfull(t *testing.T) {
	consumer := dummyConsumer{}
	broker := NewBroker(consumer)
	observable := dummyObservable{launchesCount: 0}

	for i := 0; i < 10; i++ {
		broker.Watch(&observable)
	}

	time.Sleep(time.Millisecond)
	assert.NotEqual(t, observable.launchesCount, 0, "Polling must be spawned")
	assert.Equal(t, observable.launchesCount, 1, "Polling must be spawned exactly one time")
}
