package main

import (
	"fmt"
	"strings"
)

func main() {
	sb := strings.Builder{}
	sb.WriteString("abc")
	sb.WriteString("def")
	fmt.Println(sb.String())
}
