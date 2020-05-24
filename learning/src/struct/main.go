package main

import (
	"fmt"
)

type person struct {
	name  string
	age   int8
	slice []int8
	map1  map[string]string
}

type pointer struct {
	x int
}

type react struct {
	leftUp, rightDown *pointer
}

func main() {
	p1 := person{name: "Lei", age: 28}
	fmt.Println(p1)

	p1.name = "lei"
	p1.age = 29
	p1.slice = append(p1.slice, 2)
	p1.map1 = map[string]string{}
	p1.map1["123"] = "456"
	fmt.Println(p1)

	pCopy := p1
	pCopy.name = "copy"
	fmt.Println(pCopy.name, p1.name)

	p2 := person{}
	p2.name = "name"
	fmt.Println(p2)

	p3 := new(person)
	p3.name = "new"
	fmt.Println(*p3)

	p4 := &person{}
	p4.name = "pointer"
	fmt.Println(*p4)
	fmt.Println(p4.name)

	r := react{leftUp: &pointer{10}, rightDown: &pointer{30}}
	fmt.Printf("leftUp address: %p, rightDown address: %p \n", r.leftUp, r.rightDown)
}
