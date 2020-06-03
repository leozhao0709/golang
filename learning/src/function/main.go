package main

import "fmt"

var a = firstRun()

func sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

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
