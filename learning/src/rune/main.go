package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {
	a := 'a'

	fmt.Println(a)
	fmt.Printf("rune byte is %d\n", unsafe.Sizeof(a))

	str := "中国"
	var runeCount = utf8.RuneCountInString(str)
	fmt.Println(runeCount)

	chars := []rune(str)
	fmt.Println(chars[0])
	fmt.Println(len(chars))
	fmt.Printf(".....%c", chars[1])
}
