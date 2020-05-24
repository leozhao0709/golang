package main

import (
	"encoding/json"
	"fmt"
)

type monster struct {
	Name   string `json:"name"`
	Age    int    `json:"age,omitempty"`
	Skill  string `json:"skill"`
	Ignore bool   `json:"-"`
}

func main() {
	// marshal
	monster1 := &monster{Name: "牛魔王", Age: 0, Skill: "芭蕉扇", Ignore: true}
	jsonStr, err := json.Marshal(*monster1)
	if err != nil {
		fmt.Println("json error", err)
	}
	fmt.Println(string(jsonStr))

	// unmarshal
	var monster2 *monster = &monster{}
	err = json.Unmarshal(jsonStr, monster2)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(*monster2)
}
