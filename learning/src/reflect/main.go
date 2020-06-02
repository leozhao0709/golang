package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	var stu = Student{
		Name: "tom",
		Age:  20,
	}

	reflectTest(stu)
}

func reflectTest(b interface{}) {

	rType := reflect.TypeOf(b)
	fmt.Println(rType)

	rVal := reflect.ValueOf(b)
	fmt.Println(rVal)

	iV := rVal.Interface()
	fmt.Printf("iv=%v ivType=%T\n", iV, iV)

	stu, ok := b.(Student)
	if ok {
		fmt.Println(stu)
	}
}
