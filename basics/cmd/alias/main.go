package main

import "fmt"

type ss map[string]int

func (s ss) Say() {
	fmt.Println("Hello, ", s["hello"])
}

func main() {
	s := ss{}
	s.Say()
}
