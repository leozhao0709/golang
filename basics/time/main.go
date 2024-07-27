package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	fmt.Println(tomorrow)
	fmt.Println(now.After(tomorrow))

	t2 := time.Now().Add(time.Minute * 2)
	fmt.Println(t2.After(now))
}
