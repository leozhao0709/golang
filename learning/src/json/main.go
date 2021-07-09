package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type monster struct {
	Name   string  `json:"name"`
	Age    int     `json:"age,omitempty"`
	Skill  *string `json:"skill"`
	Ignore bool    `json:"-"`
}

func main() {
	testStruct()
	// testMap()
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
	var skill *string = nil
	monster1 := &monster{Name: "牛魔王", Age: 0, Skill: skill, Ignore: true}
	jsonStr, err := json.Marshal(monster1)
	if err != nil {
		fmt.Println("json error", err)
		return
	}
	fmt.Println(string(jsonStr))

	// unmarshal
	// var monster2 = &monster{}
	var monster2 monster = monster{}
	err = json.Unmarshal(jsonStr, &monster2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(monster2)

	var monster3 = make(map[string]interface{})
	err = json.Unmarshal(jsonStr, &monster3)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(monster3)

	// Filter Int
	arr := []int{1, 2, 3, 4, 5}
	result := filter(arr, func(val interface{}) bool {
		return val != 5
	})
	fmt.Println(result)
	// [1 2 3 4]
	// filter String
	arr1 := []string{"a", "b", "c", "d", "e"}
	result1 := filter(arr1, func(val interface{}) bool {
		return val != "c"
	})
	fmt.Println(result1)
	// [a b d e]
}

func filter(arr interface{}, cond func(interface{}) bool) interface{} {
	contentType := reflect.TypeOf(arr)
	contentValue := reflect.ValueOf(arr)

	newContent := reflect.MakeSlice(contentType, 0, 0)
	for i := 0; i < contentValue.Len(); i++ {
		if content := contentValue.Index(i); cond(content.Interface()) {
			newContent = reflect.Append(newContent, content)
		}
	}
	return newContent.Interface()

	// // not working
	// arr1 := arr.([]interface{})
	// var res []interface{}
	// for _, value := range arr1 {
	// 	if cond(value) {
	// 		res = append(res, value)
	// 	}
	// }
	// return res
}
