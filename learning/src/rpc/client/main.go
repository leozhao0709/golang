package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	cli, err := rpc.DialHTTP("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("err", err)
		return
	}

	var pd int
	err = cli.Call("Panda.GetInfo", 10086, &pd)

	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("receive pd:", pd)
}
