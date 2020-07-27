package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/leozhao0709/learning/src/protobuf/person"
)

func main() {
	info := &person.PersonInfo{
		Info: []*person.PersonInfo_Person{{
			Name:   "lei",
			Height: 181,
			Weight: []int32{120, 125, 140, 160},
		}},
	}
	fmt.Println(info) // info:{name:"lei"  height:181  weight:120  weight:125  weight:140  weight:160}

	data, err := proto.Marshal(info)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data) // binary data

	info2 := &person.PersonInfo{}
	err = proto.Unmarshal(data, info2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info2) // original info2
}
