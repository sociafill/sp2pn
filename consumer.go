package main

// Consumer is an object which receive calls every portion of pulled data
type Consumer interface {
	Consume(Observable, interface{})
}
