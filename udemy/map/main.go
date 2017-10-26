package main

import (
	"fmt"
)

func main() {
	// colors := make(map[string]string)
	// colors := map[string]string{}

	colors := map[string]string{
		"red":   "ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	delete(colors, "abc")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
