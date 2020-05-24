package main

import "fmt"

type Person struct {
	Name string
	Age  int8
}

func (p *Person) speak() {
	fmt.Printf("%v is a good guy\n", p.Name)
}

func (p *Person) String() string {
	return fmt.Sprintf("Name=[%v] Age=[%d]", p.Name, p.Age)
}

func main() {
	p := &Person{Name: "Lei", Age: 28}
	p.speak()
	fmt.Println(p)
}
