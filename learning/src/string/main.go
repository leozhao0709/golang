package main

import "fmt"

func main() {
	var num1 = 99
	// var num2 = 23.456
	// var b = true
	// var myChar = 'h'

	// fmt.Sprintf
	num1Str := fmt.Sprintf("%d", num1)
	fmt.Printf("str type is %T str=%v", num1Str, num1Str)
}
