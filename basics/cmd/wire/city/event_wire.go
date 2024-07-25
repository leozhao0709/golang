//go:build wireinject
// +build wireinject

package city

import (
	"example.com/basics/cmd/wire/event"
	"github.com/google/wire"
)

// InitializeEvent creates an Event. It will error if the Event is staffed with
// a grumpy greeter.

var EventSet = wire.NewSet(event.EventSet)

func InitializeEvent(phrase string) (event.Event, error) {
	wire.Build(EventSet)
	return event.Event{}, nil
}
