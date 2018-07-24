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
	interval      time.Duration
}

func (observable *dummyObservable) Identifier() Identifier {
	return "dummy"
}

func (observable *dummyObservable) Poll() []interface{} {
	observable.launchesCount++
	var result []interface{}
	return result
}

func (observable *dummyObservable) GetInterval() time.Duration {
	return observable.interval
}

func TestNewBrokerSuccessful(t *testing.T) {
	consumer := dummyConsumer{}
	NewBroker(consumer)
}

func TestPollingIsSpawnedSuccessfully(t *testing.T) {
	consumer := dummyConsumer{}
	broker := NewBroker(consumer)
	observable := dummyObservable{launchesCount: 0, interval: time.Second}

	for i := 0; i < 10; i++ {
		broker.Watch(&observable)
	}

	time.Sleep(time.Millisecond)
	assert.NotEqual(t, 0, observable.launchesCount, "Polling must be spawned")
	assert.Equal(t, 1, observable.launchesCount, "Polling must be spawned exactly one time")
}

func TestPollingIsStoppedSuccessfully(t *testing.T) {
	consumer := dummyConsumer{}
	broker := NewBroker(consumer)
	observable := dummyObservable{launchesCount: 0, interval: 0}
	broker.Watch(&observable)
	time.Sleep(time.Millisecond)
	assert.NotEqual(t, 0, observable.launchesCount, "Polling must be spawned")
	broker.Unwatch(&observable)
	time.Sleep(time.Millisecond)
	launchesCount := observable.launchesCount
	time.Sleep(time.Millisecond)
	assert.Equal(t, launchesCount, observable.launchesCount, "Polling must be stopped")
}
