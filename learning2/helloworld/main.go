package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type address struct {
	zipCode int
}

type person struct {
	name string
	address
}

func (p *person) updateName(name string) {
	p.name = name
}

func updateName(p *person) {
	p.name = "12345"
}

func updateAddressZipCode(p person) {
	p.address = address{
		zipCode: 100,
	}
	fmt.Printf("%p", &p.address)
	fmt.Println()
}

type NameUpdate interface {
	updateName(name string)
}

func main() {

}

func useFlag() {
	var user = flag.String("u", "default user", "user description")
	var pwd = flag.String("pwd", "12345", "password description")
	var host = flag.String("h", "127.0.0.1", "host description")
	var port *int = flag.Int("p", 3306, "port description")

	flag.Parse() // must have this step

	fmt.Printf("user=%v, pwd=%v, host=%v, port=%v\n", *user, *pwd, *host, *port)
}

func HelloMakeChanSize() {
	size := 3
	c1 := make(chan int, size)
	go func() {
		for i := 0; i < 4; i++ {
			val := i*10 + 7
			fmt.Println(time.Now(), "<- ", val, "at", i)
			c1 <- i*10 + 7
		}
		c1 <- 0
	}()
	time.Sleep(time.Second * 3)
	fmt.Println("After Sleep")
	for val := range c1 {
		fmt.Println(time.Now(), "receive:", val)
		if val == 0 {
			break
		}
	}
}

func chanTest() {
	message := make(chan string, 3)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer close(message)
		message <- "hello chan 1"
		message <- "hello chan 2"
		message <- "hello chan 3"
	}()

	for {
		if result, ok := <-message; ok {
			fmt.Println(result)
		} else {
			break
		}
	}

	wg.Done()
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
			return "time out"
		}
	}
}

func service1(c chan<- string) {
	fmt.Println("service1 started!")
	rand.Seed(time.Now().UnixNano())
	resTime := rand.Intn(200)
	fmt.Println("service1 resTime is", resTime)
	time.Sleep(time.Millisecond * time.Duration(resTime))
	c <- "service1 response"
}

func service2(c chan<- string) {
	fmt.Println("service2 started!")
	rand.Seed(time.Now().UnixNano())
	resTime := rand.Intn(200)
	fmt.Println("service2 resTime is", resTime)
	time.Sleep(time.Millisecond * time.Duration(resTime))
	c <- "service1 response"
}
