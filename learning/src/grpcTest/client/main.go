package main

import (
	"context"
	"fmt"

	"github.com/leozhao0709/learning/src/grpcTest/proto/hello"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:8000", grpc.WithInsecure())
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer conn.Close()

	client := hello.NewHelloServerClient(conn)

	helloRes, err := client.SayHello(context.Background(), &hello.HelloReq{Name: "panda"})
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("SayHello Response", helloRes)

	nameRes, err := client.SayName(context.Background(), &hello.NameReq{Name: "托尼斯塔克"})
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("SayHello Response", nameRes)
}
