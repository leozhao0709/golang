package main

import "fmt"

func main() {
	a := []int{}
	a = append(a, 1, 2, 3)

	appendTest(a, 1)
	fmt.Println("------", a)

	appendTest(a, 10)
	fmt.Println("------", a)

	b := a[2:3]
	fmt.Println(b)
}

func appendTest(arr []int, n int) {
	arr[0] = 100
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}
}
