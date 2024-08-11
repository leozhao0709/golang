package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("select----", useSelect())
}

func useSelect() string {

	fmt.Println("useSelecte started!")

	var chan1 chan string = make(chan string)
	var chan2 chan string = make(chan string)

	go service1(chan1)
	go service2(chan2)

	timeOutChan := time.After(time.Millisecond * 100)
	for {
		select {
		case <-chan1:
			// fmt.Println("response from service1")
			return "response from chan1"
		case <-chan2:
			// fmt.Println("response from service2")
			return "response from chan2"
		case <-time.After(time.Millisecond * 50):
			fmt.Println("more than 50ms response")
		case <-timeOutChan:
			// fmt.Println("...time out...")
			return "...time out... with no response from either service"
		}
	}
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
