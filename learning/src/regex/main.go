package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	match, _ := regexp.MatchString("p(\\w+)ch", "peach")
	fmt.Println(match) // true

	r, _ := regexp.Compile("p(\\w+)ch")
	fmt.Println(r.MatchString("peach")) // true

	fmt.Println(r.FindString("peach punch"))        // peach, only find 1 time
	fmt.Println(r.FindAllString("peach punch", 1))  // [peach], only find 1 time
	fmt.Println(r.FindAllString("peach punch", -1)) // [peach punch], -1 will find all

	fmt.Println(r.FindStringSubmatch("peach punch"))        // [peach ea] only find 1 time
	fmt.Println(r.FindAllStringSubmatch("peach punch", -1)) // [[peach ea] [punch un]]

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r) // p([a-z]+)ch

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))           // a <fruit>
	fmt.Println(r.ReplaceAllStringFunc("a peach", strings.ToUpper)) // a PEACH
}
