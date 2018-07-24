package sp2pt

import (
	"time"
)

type identifiersMap map[Identifier]int
type identifiersControl map[Identifier]chan bool

// Broker is main structure which allows to manage observables, subscriptions, consumer
type Broker struct {
	identifiers identifiersMap
	observables identifiersControl
	consumer    Consumer
}

// NewBroker is constructor for Broker
func NewBroker(consumer Consumer) Broker {
	broker := Broker{}
	broker.identifiers = identifiersMap{}
	broker.observables = identifiersControl{}
	broker.consumer = consumer
	return broker
}

// Watch function is used to initiate watching to some observable object
func (broker *Broker) Watch(observable Observable) {
	broker.identifiers[observable.Identifier()]++
	if !broker.isRunning(observable) {
		broker.runPolling(observable)
	}
}

// Unwatch function is used to decrement watchers counter for some observable object.
// If some observable object has no watchers - polling must be stopped
func (broker *Broker) Unwatch(observable Observable) {
	broker.identifiers[observable.Identifier()]--
	if broker.identifiers[observable.Identifier()] < 1 {
		if broker.isRunning(observable) {
			broker.stopPolling(observable)
		}
	}
}

func (broker *Broker) isRunning(observable Observable) bool {
	_, isRunning := broker.observables[observable.Identifier()]
	return isRunning
}

func (broker *Broker) runPolling(observable Observable) {
	channel := make(chan bool)
	broker.observables[observable.Identifier()] = channel
	go polling(channel, observable, broker.consumer)
}

func (broker *Broker) stopPolling(observable Observable) {
	channel := broker.observables[observable.Identifier()]
	close(channel)
}

func polling(stop chan bool, observable Observable, consumer Consumer) {
	for {
		select {
		default:
			items := observable.Poll()
			for _, item := range items {
				consumer.Consume(observable, item)
			}
		case <-stop:
			return
		}
		time.Sleep(observable.GetInterval())
	}
}
