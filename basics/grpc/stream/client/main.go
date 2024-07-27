package main

import (
	"context"
	"io"
	"log"
	"strconv"
	"sync"
	"time"

	pb "example.com/basics/grpc/stream/protogen/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// getServerSendStream(client)
	// pushStream(client)
	bidirectionalStream(client)
}

func getServerSendStream(client pb.GreeterClient) {
	ctx := context.Background()
	stream, err := client.ServerSendStream(ctx, &pb.StreamReqData{Data: "hello"})

	if err != nil {
		log.Fatalf("Failed to get stream: %v", err)
	}

	for {
		streamData, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				log.Fatalf("error receiving data: %v", err)
			}
			break
		}
		log.Println(streamData.Data)
	}

	log.Println("stream closed")
}

func pushStream(client pb.GreeterClient) error {
	stream, err := client.ClientSendStream(context.Background())
	if err != nil {
		log.Fatalf("Failed to get stream: %v", err)
		return err
	}
	for i := 0; i < 5; i++ {
		stream.Send(&pb.StreamReqData{Data: "world " + strconv.Itoa(i)})
		time.Sleep(time.Second)
	}
	return nil
}

func bidirectionalStream(client pb.GreeterClient) error {
	stream, err := client.BidirectionalStream(context.Background())
	if err != nil {
		log.Fatalf("Failed to get stream: %v", err)
		return err
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					log.Printf("error receiving data: %v", err)
				}
				break
			}
			log.Println("client received: ", data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			stream.Send(&pb.StreamReqData{Data: "hello from client " + strconv.Itoa(i)})
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	return nil
}
