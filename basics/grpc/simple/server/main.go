package main

import (
	"context"
	"log"
	"net"

	pb "example.com/basics/grpc/simple/gen/go/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// Ping implements pb.GreeterServer.
func (s *server) Ping(context.Context, *emptypb.Empty) (*pb.Pong, error) {
	return &pb.Pong{Message: "pong"}, nil
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello " + in.GetName(),
	}, nil
}

var _ pb.GreeterServer = (*server)(nil)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}