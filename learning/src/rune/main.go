package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 'a'

	fmt.Println(a)
	fmt.Printf("rune byte is %d\n", unsafe.Sizeof(a))

	str := "中国"
	chars := []rune(str)
	fmt.Println(chars[0])
	fmt.Printf(".....%c", chars[1])
}
