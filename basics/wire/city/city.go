package city

import "example.com/basics/wire/event"

type City struct {
	Event event.Event
}

func NewCity(e event.Event) City {
	return City{Event: e}
}
