package main

import (
	"fmt"
	"os"

	pb "example.com/basics/grpc/simple/protogen/v1"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	hi := pb.HiResponse{
		Message: "Hello, World!",
		Status:  pb.Status_OK,
	}

	marshaler1 := jsonpb.Marshaler{EmitDefaults: true}
	marshaler1.Marshal(os.Stdout, &hi)

	marshaler2 := protojson.MarshalOptions{EmitUnpopulated: true}
	res, err := marshaler2.Marshal(&hi)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	unmarshaler := protojson.UnmarshalOptions{}
	hi2 := pb.HiResponse{}
	err = unmarshaler.Unmarshal([]byte(`{"message":"Hello, World!"}`), &hi2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", &hi2)

	hi3 := pb.HiRequest{}
	err = unmarshaler.Unmarshal([]byte(`{"message":"Hello, World!"}`), &hi3)

	if err != nil {
		panic(err)
	}
	fmt.Print(hi3.GetName())
}
