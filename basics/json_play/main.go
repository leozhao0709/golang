package main

import (
	"encoding/json"
	"fmt"
)

type PersonReq1 struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

func main() {
	p1 := &PersonReq1{
		Email: "john@example.com",
		Name:  "John Doe",
		Age:   0,
	}

	bytes, _ := json.Marshal(p1)
	fmt.Println(string(bytes))

	str := `{"email":"john@example.com","name":"John Doe"}`
	p2 := &PersonReq1{}
	err := json.Unmarshal([]byte(str), &p2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)
}
