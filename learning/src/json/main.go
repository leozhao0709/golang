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
	// testStruct()
	testMap()
}

func testMap() {
	// marshal
	var a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "address"

	jsonStr, err := json.Marshal(a)
	if err != nil {
		fmt.Println("marshal error", err)
		return
	}
	fmt.Println(string(jsonStr))

	// unmarshal
	var a1 map[string]interface{}
	err = json.Unmarshal(jsonStr, &a1)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(a1)
}

func testStruct() {
	// marshal
	monster1 := &monster{Name: "牛魔王", Age: 0, Skill: "芭蕉扇", Ignore: true}
	jsonStr, err := json.Marshal(monster1)
	if err != nil {
		fmt.Println("json error", err)
		return
	}
	fmt.Println(string(jsonStr))

	// unmarshal
	// var monster2 = &monster{}
	var monster2 *monster = &monster{}
	err = json.Unmarshal(jsonStr, monster2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(*monster2)
	fmt.Println(monster2.Name)
}
