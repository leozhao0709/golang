package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

// Panda ...
type Panda int

// GetInfo ...
func (panda *Panda) GetInfo(argType *int, replyType *int) error {
	fmt.Println("client send", *argType)

	*replyType = *argType + 1
	return nil
}

func main() {
	// 1. we need initial a object
	pd := new(Panda)

	// 2. register object
	rpc.Register(pd)
	rpc.HandleHTTP()

	err := http.ListenAndServe("0.0.0.0:8080", nil)

	if err != nil {
		fmt.Println("err", err)
		return
	}
}
