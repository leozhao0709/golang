package main

import (
	"fmt"
)

func main() {
	test :=
		`姓名		年龄		籍贯		住址
	john		12		河北		北京`
	fmt.Println(test)

	test1 := "abc\ndef"
	fmt.Println(test1)

	test2 := `abc\ndef`
	fmt.Println(test2)

	test3 := "this is a " +
		"long string " +
		"in multiple line"
	fmt.Println(test3)
}
