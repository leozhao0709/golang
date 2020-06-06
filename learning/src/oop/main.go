package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name string
	Age  int
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
	b.Age = 18

	fmt.Println("b.age is", b.Age)
	b.SayOk()
	b.A.SayOk()
	b.hello()

	b2 := B{Name: "b2", A: A{Name: "a", Age: 18}}
	fmt.Printf("%+v\n", b2)

	jsonStr, _ := json.Marshal(&b2)
	fmt.Println(string(jsonStr))
}
