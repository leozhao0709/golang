package main

import (
	"fmt"
	"sync"
)

func f1(ch chan int) {
	// defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func f2(ch1 chan int, ch2 chan int) {
	// defer wg.Done()
	for x := range ch1 {
		ch2 <- x * x
	}
	close(ch2)
}

var wg sync.WaitGroup

func main() {

	a := make(chan int)
	b := make(chan int)

	// wg.Add(2)
	go f1(a)
	go f2(a, b)

	// wg.Wait()
	for x := range b {
		fmt.Println(x)
	}

	// a <- 1
	// close(a)
	// for x := range a {
	// 	fmt.Println(x)
	// }

	// useChannel()
	// useWg()
	// useSelect()
	// useMutex()
}
