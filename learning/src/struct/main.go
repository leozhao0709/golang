package main

import "fmt"

type person struct {
	name string
	age  int8
}

func main() {
	p := person{name: "Lei", age: 28}
	fmt.Println(p)

	p.name = "lei"
	p.age = 29
	fmt.Println(p)
}
