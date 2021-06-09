package main

import "fmt"

type Point struct {
	x, y int
}

func main() {
	var a interface{}
	var point = Point{1, 2}
	a = point

	var b Point
	b, ok := a.(Point)

	if ok {
		fmt.Println(b)
	} else {
		fmt.Println("convert failed....")
	}

	c := a.(string)
	fmt.Println(c)
}
