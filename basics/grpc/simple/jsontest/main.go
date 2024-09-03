package main

import (
	"os"

	pb "example.com/basics/grpc/simple/protogen/v1"
	"github.com/golang/protobuf/jsonpb"
)

func main() {
	hi := pb.HiResponse{
		Message: "Hello, World!",
		Status:  pb.Status_OK,
	}

	marshaler := jsonpb.Marshaler{EmitDefaults: true}
	marshaler.Marshal(os.Stdout, &hi)
}
