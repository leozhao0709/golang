package main

import (
	"io"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	pb "example.com/basics/grpc/stream/protogen/v1"
	"google.golang.org/grpc"
)

type streamServer struct {
	pb.UnimplementedGreeterServer
}

func (s *streamServer) ServerSendStream(in *pb.StreamReqData, stream pb.Greeter_ServerSendStreamServer) error {
	for i := 0; i < 5; i++ {
		err := stream.Send(&pb.StreamResData{Data: in.Data + " " + strconv.Itoa(i)})
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (s *streamServer) ClientSendStream(stream pb.Greeter_ClientSendStreamServer) error {
	for {
		data, err := stream.Recv()
		if err != nil {
			log.Printf("error receiving data: %v", err)
			break
		}
		log.Println(data.Data)
	}

	log.Println("stream closed")
	return nil
}

func (s *streamServer) BidirectionalStream(stream pb.Greeter_BidirectionalStreamServer) error {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			streamData, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					log.Printf("error receiving data: %v", err)
				}
				break
			}
			log.Println("server received: ", streamData.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			stream.Send(&pb.StreamResData{Data: "hello from server " + strconv.Itoa(i)})
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	return nil
}

var _ pb.GreeterServer = (*streamServer)(nil)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &streamServer{})

	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
