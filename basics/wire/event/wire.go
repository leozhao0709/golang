//go:build wireinject
// +build wireinject

package event

import "github.com/google/wire"

// InitializeEvent creates an Event. It will error if the Event is staffed with
// a grumpy greeter.

var EventSet = wire.NewSet(NewEvent, NewGreeter, NewMessage)

func InitializeEvent(phrase string) (Event, error) {
	wire.Build(EventSet)
	return Event{}, nil
}
