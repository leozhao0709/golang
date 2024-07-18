package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	a := make([]int, 0)
	a = append(a, 1, 2, 3)
	fmt.Println(a)
	b := a[:2]
	fmt.Println(b)

	var c []int
	// c = append(c, 2)
	fmt.Println(c)

	rand := rand.New(rand.NewSource(66))
	fmt.Println(rand.Intn(100))

	d := [...]int{1, 2, 3}
	s := strings.Builder{}
	for _, ele := range d {
		s.WriteString(strconv.Itoa(ele))
	}

	fmt.Println(s.String())
}
