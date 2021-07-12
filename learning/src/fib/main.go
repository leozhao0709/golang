package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// var resChan = make(chan int, 0)

	var testVar = 47
	// go func() {
	// 	resChan <- fib(testVar)
	// }()

	var res = fib(testVar)

	// var resultFromChan = <-resChan
	// fmt.Println("result from go routine is ", resultFromChan)
	fmt.Println("result from main is ", res)

	duration := time.Since(start)
	fmt.Println("running duration", duration)
}

func fib(n int) int {
	if n == 1 || n == 0 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}
