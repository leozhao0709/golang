package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("goroutine helloworld! %d\n", i)
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Printf("main() helloworld! %d\n", i)
		time.Sleep(time.Second)
	}
}

func setCPUNum() {
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)

	// set how many cpy core are using
	runtime.GOMAXPROCS(cpuNum)
}
