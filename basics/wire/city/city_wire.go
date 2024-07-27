//go:build wireinject
// +build wireinject

package city

import (
	"example.com/basics/wire/event"
	"github.com/google/wire"
)

var CitySet = wire.NewSet(NewCity, event.EventSet)

func InitializeCity(eventName string) (City, error) {
	wire.Build(CitySet)
	return City{}, nil
}
