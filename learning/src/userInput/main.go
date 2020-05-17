package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Println("what's your name?")
	fmt.Scanln(&name)
	fmt.Println("How old are you?")
	fmt.Scanln(&age)
	fmt.Printf("%s is %d years old \n", name, age)
}
