package main

import "fmt"

func useChannel() {
	var intchan = make(chan int)
	var readDoneChan = make(chan bool)
	var writeDoneChan = make(chan bool)

	go writeDataWithChannel(intchan, writeDoneChan)
	go readDataWithChannel(intchan, readDoneChan)
	<-writeDoneChan
	<-readDoneChan
}

func writeDataWithChannel(intchan chan int, writeDoneChan chan bool) {
	for i := 0; i < 10; i++ {
		intchan <- i
		fmt.Println("writeData", i)
	}
	fmt.Println("write data end...")
	writeDoneChan <- true
	close(intchan)
}

func readDataWithChannel(intChan chan int, doneChan chan bool) {
	for value := range intChan {
		fmt.Println("readData", value)
	}

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("readData", <-intChan)
	// }

	fmt.Println("read data end...")
	doneChan <- true
}
