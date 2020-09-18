package main

import (
	"fmt"
	"time"
)

func main() {
	// timeAfter()
	timeTicker()
}

func newTimer() {
	fmt.Println("current time1", time.Now())
	myTimer := time.NewTimer(time.Second * 2)
	<-myTimer.C
	fmt.Println("current time2", time.Now())
}

func timeAfter() {
	fmt.Println("current time1", time.Now())
	<-time.After(time.Second * 2)
	fmt.Println("current time2", time.Now())
}

func timeTicker() {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case now := <-ticker.C:
			fmt.Println("currentTime:", now)
		default:
		}
	}
}
