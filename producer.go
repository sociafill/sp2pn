package main

type Producer interface {
	Identifier() Identifier
	Poll() []interface{}
}
