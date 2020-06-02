package main

import (
	"fmt"
	"sync"
)

var i = 0

func worker(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	i++
	wg.Done()
}

func useMutex() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &mutex)
	}

	wg.Wait()

	fmt.Println("value of i after 1000 operations is", i)
}
