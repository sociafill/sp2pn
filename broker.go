package main

type Broker struct {
	identifiers map[Identifier]int
	producers   map[Identifier]chan bool
	consumer    Consumer
}

func NewBroker(consumer Consumer) Broker {
	broker := Broker{}
	broker.identifiers = map[Identifier]int{}
	broker.producers = map[Identifier]chan bool{}
	broker.consumer = consumer
	return broker
}
