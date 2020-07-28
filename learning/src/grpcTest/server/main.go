package main

import (
	"context"
	"fmt"
	"net"

	"github.com/leozhao0709/learning/src/grpcTest/proto/hello"
	"google.golang.org/grpc"
)

// Student ...
type Student struct {
}

// SayHello ...
func (s *Student) SayHello(ctx context.Context, in *hello.HelloReq) (*hello.HelloRes, error) {
	return &hello.HelloRes{
		Msg: "hello," + in.Name,
	}, nil
}

// SayName ...
func (s *Student) SayName(ctx context.Context, in *hello.NameReq) (*hello.NameRes, error) {
	return &hello.NameRes{
		Msg: "Name," + in.Name,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println("err", err)
		return
	}

	// create grpc server
	gserver := grpc.NewServer()
	// register service
	hello.RegisterHelloServerServer(gserver, &Student{})

	err = gserver.Serve(listener)
	if err != nil {
		fmt.Println("err", err)
		return
	}
}
