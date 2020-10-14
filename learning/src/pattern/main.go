package main

import (
	"fmt"

	"github.com/leozhao0709/learning/src/pattern/user"
)

func main() {
	var srv1 = user.Service
	var srv2 = user.Service

	fmt.Printf("...srv1 %p", srv1)
	fmt.Printf("...srv2 %p", srv2)

	var age1 = srv1.GetAge()
	fmt.Println(age1)
}
