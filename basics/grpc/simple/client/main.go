package main

import (
	"context"
	"log"
	"time"

	pb "example.com/basics/grpc/simple/gen/go/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	conn, err := grpc.NewClient("0.0.0.0:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	respPong, err := client.Ping(ctx, &emptypb.Empty{})
	respHello, err := client.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Greeted: %s", respPong.GetMessage())
	log.Printf("Greeted: %s", respHello.GetMessage())
}
