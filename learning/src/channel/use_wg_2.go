package main

import (
	"fmt"
)

func useWg2() {

	// var wg sync.WaitGroup

	a := make(chan int)
	// b := make(chan int, 10)
	b := make(chan int)

	// wg.Add(2)
	go func(ch chan int) {
		// defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("write.....", i)
			ch <- i
		}
		close(ch)
	}(a)

	go func(ch1 chan int, ch2 chan int) {
		// defer wg.Done()
		for x := range ch1 {
			fmt.Println("....calculate..", x, "....", x*x)
			ch2 <- x * x
		}
		close(ch2)
	}(a, b)

	// wg.Wait()
	for x := range b { // block until channel b close
		fmt.Println(x)
	}

	// a <- 1
	// close(a)
	// for x := range a {
	// 	fmt.Println(x)
	// }
}
