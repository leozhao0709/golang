package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func useAtomic() {
	var x int64 = 0

	var wg sync.WaitGroup

	add := func() {
		atomic.AddInt64(&x, 1)
		wg.Done()
	}

	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go add()
	}

	wg.Wait()
	fmt.Println(x)
}
