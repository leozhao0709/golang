package main

import "fmt"

func main() {
	// var intChan chan int = make(chan int, 3)
	// intChan <- 10
	// num := 211
	// intChan <- num
	// intChan <- 50

	// fmt.Printf("intChan len=%d cap=%d\n", len(intChan), cap(intChan))

	// num2 := <-intChan
	// fmt.Println("num2 is", num2)
	// close(intChan)

	iterateChan()
}

func iterateChan() {
	var intChan2 chan int = make(chan int, 100)

	for i := 0; i < cap(intChan2); i++ {
		intChan2 <- i * 2
	}

	// must close channel before you iterate channel
	close(intChan2)

	// must use for range channel loop to iterate channel
	for value := range intChan2 {
		fmt.Println(value)
	}
}
