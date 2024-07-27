package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := 5
	// a_bin := fmt.Sprintf("%b", a)
	a_bin := strconv.FormatInt(int64(a), 2)
	fmt.Println(a_bin)

	fmt.Println(((-2 % 3) + 3) % 3)

	fmt.Println(strings.Trim("###asdasd  %", "#% "))

	b := fmt.Sprintf("%5d", a)
	fmt.Println(b)
}
