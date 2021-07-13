package main

import (
	"context"
	"fmt"
	"sync"
)

var (
	ticketsChan = make(chan int, 1)
	wg          sync.WaitGroup
	mutex       sync.Mutex
)

func main() {
	ticketsAmount := 10000
	ticketsChan <- ticketsAmount

	for i := 0; i < 3; i++ {
		wg.Add(1)
		ctx := context.WithValue(context.Background(), "id", fmt.Sprintf("goroutine %v", i))
		go sellTicket(ctx)
	}

	wg.Wait()
}

func sellTicket(ctx context.Context) {
	defer wg.Done()
	for {
		leftTicket, ok := <-ticketsChan
		if ok {
			leftTicket--
			fmt.Println(ctx.Value("id"), ": Sell ticket left", leftTicket)
			if leftTicket == 0 {
				fmt.Println(ctx.Value("id"), ": Sell complete")
				close(ticketsChan)
				return
			}
			ticketsChan <- leftTicket
		} else {
			fmt.Println(ctx.Value("id"), ": Sell already complete before")
			return
		}
	}
}
