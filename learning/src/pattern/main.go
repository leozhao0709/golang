package main

import (
	"fmt"

	"github.com/leozhao0709/learning/src/pattern/user"
)

func main() {
	srv1 := user.Service
	srv2 := user.Service

	fmt.Printf("...srv1 %p", srv1)
	fmt.Printf("...srv2 %p", srv2)
}
