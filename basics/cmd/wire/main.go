package main

import (
	"fmt"
	"os"

	"example.com/basics/cmd/wire/event"
)

func main() {
	e, err := event.InitializeEvent("hi there!")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}
