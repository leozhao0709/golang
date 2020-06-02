package main

import (
	"fmt"
	"sync"
)

func useWg() {
	fmt.Println("use wg start...")
	var wg sync.WaitGroup
	var intchan = make(chan int)

	wg.Add(1)
	go writeDataWithWg(intchan)
	go readDataWithWg(intchan, &wg)

	wg.Wait()
	fmt.Println("use wg stop...")
}

func writeDataWithWg(intchan chan int) {
	for i := 0; i < 10; i++ {
		intchan <- i
		fmt.Println("writeData", i)
	}
	fmt.Println("write data end...")
	close(intchan)
}

func readDataWithWg(intchan chan int, wg *sync.WaitGroup) {
	for value := range intchan {
		fmt.Println("readData", value)
	}

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("readData", <-intChan)
	// }

	fmt.Println("read data end...")
	wg.Done()
}
