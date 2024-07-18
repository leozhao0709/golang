package main

import (
	"fmt"
)

func main() {
	s := "中，Hello"
	fmt.Println(len([]rune(s)))
	r := []rune(s)
	r[0] = '哈'
	fmt.Println(string(r))
}
