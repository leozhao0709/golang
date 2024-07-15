package string

import "fmt"

func ExampleReverse() {
	result := Reverse("abcde")
	fmt.Println(result)
	// Output:
	// edcba
}
