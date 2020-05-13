package main

import (
	"fmt"
	"unsafe"
)

var (
	a = "a"
	b = "b"
)

func main() {
	c, d := "c", 4
	fmt.Println(a, b, c, d)
	var a uint8 = 1
	fmt.Printf("bit of test is %d", unsafe.Sizeof(a))
}
