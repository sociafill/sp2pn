package main

type Consumer interface {
	Consume(interface{})
}
