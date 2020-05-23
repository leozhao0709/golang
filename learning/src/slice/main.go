package main

import "fmt"

func main() {
	var intArr = [...]int{11, 22, 33, 44, 55, 66, 99}
	slice := intArr[:5]

	// intArr[2] = 333
	// fmt.Printf("intArr 2nd position %p\n", &intArr[1])

	// fmt.Println(slice)
	// fmt.Printf("slice 1nd position %p\n", &slice[0])
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

	// slice[0] = 222
	// fmt.Println(slice)
	// fmt.Println(intArr)

	slice = append(slice[:3], slice[4:]...)
	fmt.Println(slice)

	slice = []int{}
	slice = append(slice, 1)
	fmt.Println(slice)
}
