package main

import (
	"fmt"
	"sync"
)

var (
	x      = 0
	rwLock sync.RWMutex
	wg     sync.WaitGroup
)

func useRwmutex() {

	read := func() {
		defer wg.Done()
		rwLock.RLock()
		fmt.Println(x)
		rwLock.RUnlock()
	}

	write := func() {
		defer wg.Done()
		rwLock.Lock()
		x++
		rwLock.Unlock()
	}

	count := 1000
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < count; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
}
