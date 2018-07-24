package main

import (
	"fmt"
	"time"
)

// Broker is main structure which allows to manage observables, subscriptions, consumer
type Broker struct {
	identifiers map[Identifier]int
	observables map[Identifier]chan bool
	consumer    Consumer
}

// NewBroker is constructor for Broker
func NewBroker(consumer Consumer) Broker {
	broker := Broker{}
	broker.identifiers = map[Identifier]int{}
	broker.observables = map[Identifier]chan bool{}
	broker.consumer = consumer
	return broker
}

// Watch function is used to initiate watching to some observable object
func (broker *Broker) Watch(observable Observable) {
	broker.identifiers[observable.Identifier()]++
	_, isRunning := broker.observables[observable.Identifier()]
	if !isRunning {
		broker.runPolling(observable)
	}
}

// Unwatch function is used to decrement watchers counter for some observable object.
// If some observable object has no watchers - polling must be stopped
func (broker *Broker) Unwatch(observable Observable) {
	fmt.Printf("Stop watching %s\n", observable.Identifier())
	_, isRunning := broker.observables[observable.Identifier()]
	if !isRunning {
		return
	}
	broker.identifiers[observable.Identifier()]--
	if broker.identifiers[observable.Identifier()] == 0 {
		channel := broker.observables[observable.Identifier()]
		close(channel)
	}
}

func (broker *Broker) runPolling(observable Observable) {
	fmt.Printf("Run watcher for %s\n", observable.Identifier())
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
			// fmt.Printf("Poll of [%s] has new iteration\n", observable.Identifier())
			items := observable.Poll()
			// fmt.Printf("Items received\n")
			for _, item := range items {
				consumer.Consume(observable, item)
			}
		case <-stop:
			fmt.Printf("Signal stop received for watcher of [%v]\n", observable.Identifier())
			return
		}
		time.Sleep(observable.GetInterval())
	}
}
