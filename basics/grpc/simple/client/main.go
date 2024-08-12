package main

import (
	"context"
	"log"
	"time"

	pb "example.com/basics/grpc/simple/protogen/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	var interceptor grpc.UnaryClientInterceptor = func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()

		log.Printf("Start Sending request: %s", method)
		err := invoker(ctx, method, req, reply, cc, opts...)
		timeTaken := time.Since(start)
		log.Printf("Received response: %s, took: %s", method, timeTaken)
		return err
	}

	// interceptor
	var interceptorOption = grpc.WithUnaryInterceptor(interceptor)
	options := []grpc.DialOption{interceptorOption, grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient("0.0.0.0:8000", options...)

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	md := metadata.New(map[string]string{
		// "Authorization": "Bearer YOUR_TOKEN_HERE", // key will be converted to lower case.
		"x-request-id": "12345", // Replace with your actual request ID.
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	respPong, err := client.Ping(ctx, &emptypb.Empty{})
	// respHello, err := client.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			log.Fatalf("Error: %v", err)
		}
		log.Printf("Error: %v, code: %d, message: %s", err, st.Code(), st.Message())
	}

	log.Printf("Greeted: %s", respPong.GetMessage())
	// log.Printf("Greeted: %s", respHello.GetMessage())
}
