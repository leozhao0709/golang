package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
	Name string
}

func (b *B) SayOk() {
	fmt.Println("B SayOk", b.Name)
}

func main() {
	var b = B{}
	b.Name = "b"
	b.A.Name = "a"
	b.age = 18

	fmt.Println("b.age is", b.age)
	b.SayOk()
	b.A.SayOk()
	b.hello()
}
