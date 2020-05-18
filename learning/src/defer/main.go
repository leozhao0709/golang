package main

import "fmt"

func deferTest(n1 *int, n2 *int) int {
	defer fmt.Println("defer 1....", *n1, *n2)

	*n1++
	*n2++

	defer fmt.Println("defer 2....", *n1, *n2)
	result := *n1 + *n2
	fmt.Println("defer Test result", result)
	return result
}

func main() {
	n1 := 10
	n2 := 10
	var sum = deferTest(&n1, &n2)
	fmt.Println("main sum is", sum)
}
