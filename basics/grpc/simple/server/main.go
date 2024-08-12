package main

import (
	"context"
	"log"
	"net"

	pb "example.com/basics/grpc/simple/protogen/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if len(md.Get("authorization")) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "missing authorization token")
		}
		log.Printf("Received token: %s", md.Get("authorization")[0]) // key will be converted to lower case
		log.Printf("Received token: %s", md.Get("x-request-id")[0])
	} else {
		return nil, status.Errorf(codes.Internal, "missing authorization token")
	}

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

	var interceptor grpc.UnaryServerInterceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Printf("Received request: %#v", info.FullMethod)
		resp, err := handler(ctx, req)
		return resp, err
	}
	interceptorOption := grpc.UnaryInterceptor(interceptor)

	s := grpc.NewServer(interceptorOption)
	defer s.GracefulStop()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
