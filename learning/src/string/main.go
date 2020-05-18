package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertWithSprintf() {
	var num1 = 99
	var num2 = 23.456
	var b = true
	var myChar = 'h'

	// fmt.Sprintf
	num1Str := fmt.Sprintf("%d", num1)
	fmt.Printf("str type is %T str=%q\n", num1Str, num1Str)

	num2Str := fmt.Sprintf("%f", num2)
	fmt.Printf("str type is %T str=%q\n", num2Str, num2Str)

	boolStr := fmt.Sprintf("%t", b)
	fmt.Printf("str type is %T str=%q\n", boolStr, boolStr)

	runeStr := fmt.Sprintf("%c", myChar)
	fmt.Printf("str type is %T str=%q\n", runeStr, runeStr)
}

func convertWithStrConv() {
	var num1 = 99
	var num2 = 23.456
	var b = true
	var myChar = 'h'

	num1Str := strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str type is %T str=%q\n", num1Str, num1Str)

	num1StrIota := strconv.Itoa(num1)
	fmt.Printf("str type is %T str=%q\n", num1StrIota, num1StrIota)

	num2Str := strconv.FormatFloat(num2, 'f', 10, 64)
	fmt.Printf("str type is %T str=%q\n", num2Str, num2Str)

	boolStr := strconv.FormatBool(b)
	fmt.Printf("str type is %T str=%q\n", boolStr, boolStr)

	runeStr := string(myChar)
	fmt.Printf("str type is %T str=%q\n", runeStr, runeStr)
}

func parseString() {
	var str = "true"
	var b bool

	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type is %T, b=%v\n", b, b)

	str = "99"
	var num1 int64
	num1, _ = strconv.ParseInt(str, 10, 64)
	fmt.Printf("num1 type is %T, num1=%v\n", num1, num1)

	var num2 int
	num2, _ = strconv.Atoi(str)
	fmt.Printf("num2 type is %T, num2=%v\n", num2, num2)

	var num3 float64
	str = "23.456"
	num3, _ = strconv.ParseFloat(str, 64)
	fmt.Printf("num3 type is %T, num3=%v\n", num3, num3)

	str = "h"
	var char []rune
	char = []rune(str)
	fmt.Printf("char type is %T, char=%c\n", char, char)

	str = "hello"
	test, err := strconv.ParseInt(str, 10, 64)
	fmt.Println(err != nil)
	fmt.Println(test)
}

func main() {
	// convertWithSprintf()
	// convertWithStrConv()
	// parseString()

	str := "hello北"
	length := len(str)
	fmt.Println("....str length is", length)

	r := []rune(str)
	fmt.Println("...rune length is", len(r))

	strArr := strings.Split("hello, world, ok", ",")
	fmt.Println(strArr[1])
}
