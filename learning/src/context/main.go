package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func f1(ctx context.Context) {
	defer wg.Done()
	for {
		fmt.Println("f1....")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

var wg sync.WaitGroup

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go f1(ctx)
	time.Sleep(time.Second * 5)
	wg.Wait()
}
