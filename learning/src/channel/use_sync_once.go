package main

import (
	"fmt"
	"sync"
)

func useSyncOnce() {
	var doOnce sync.Once
	doSomething := func() {
		fmt.Println("...run every time2...")
		doOnce.Do(
			func() {
				fmt.Println("....only run once...")
			})
	}

	go doSomething()
	go doSomething()
}
