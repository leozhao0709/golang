package main

import "log"

type Person struct {
	name string
}

func main() {
	p1 := Person{"Alice"}
	p2 := p1
	p2.name = "Bob"
	log.Println(p1.name) // Output: Alice
	log.Println(p2.name) // Output: Bob
}
