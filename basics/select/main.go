package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("select----", useSelect())

	// selectOverTest()
}

func selectOverTest() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1: Send a message after 2 seconds
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from ch1"
		fmt.Println("Goroutine 1 finished")
	}()

	// Goroutine 2: Send a message after 3 seconds
	go func() {
		go func() {
			time.Sleep(2 * time.Second)
			ch2 <- "Message from ch2"
		}()
		select {
		case msg := <-ch2:
			fmt.Println("in goroutine", msg)
		case <-time.After(3 * time.Second):
			fmt.Println("After 3s, No message received in ch2")
		}

		fmt.Println("Goroutine 2 finished")
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg) // This will likely execute first since ch1 sends after 2 seconds
	case msg := <-ch2:
		fmt.Println(msg) // This may execute if the timing aligns perfectly
	}

	// Sleep to ensure we see the output from the second goroutine if it's still running
	time.Sleep(4 * time.Second)
}

func useSelect() string {

	fmt.Println("useSelecte started!")

	var chan1 chan string = make(chan string)
	var chan2 chan string = make(chan string)

	go service1(chan1)
	go service2(chan2)

	timeOutChan := time.After(time.Millisecond * 100)
	// for {
	select {
	case <-chan1:
		// fmt.Println("response from service1")
		return "response from chan1"
	case <-chan2:
		// fmt.Println("response from service2")
		return "response from chan2"
	case <-time.After(time.Millisecond * 50):
		// fmt.Println("more than 50ms response")
		return "more than 50ms response"
	case <-timeOutChan:
		// fmt.Println("...time out...")
		return "...time out... with no response from either service"
	}
	// }
}

func service1(c chan<- string) {
	fmt.Println("service1 started!")

	randSrc := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(randSrc)
	resTime := rand.Intn(200)
	fmt.Println("service1 resTime is", resTime)
	time.Sleep(time.Millisecond * time.Duration(resTime))
	c <- "service1 response"
}

func service2(c chan<- string) {
	fmt.Println("service2 started!")
	randSrc := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(randSrc)
	resTime := rand.Intn(200)
	fmt.Println("service2 resTime is", resTime)
	time.Sleep(time.Millisecond * time.Duration(resTime))
	c <- "service1 response"
}
