package main

import "fmt"

var a = firstRun()

func firstRun() int {
	fmt.Println("...firstRun...")
	return 90
}

func init() {
	fmt.Println("...init()...")
}

func main() {
	fmt.Println("...main()...")
}
