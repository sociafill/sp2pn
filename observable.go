package sp2pt

import "time"

// Observable must be implemented by object which should be polled
type Observable interface {
	Identifier() Identifier
	Poll() []interface{}
	GetInterval() time.Duration
}
