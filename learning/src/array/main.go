package main

import "fmt"

func main() {
	var numsArr1 = [3]int{1, 2, 3}
	var numsArr2 = [...]int{4, 5, 6}
	var names = [3]string{1: "tom", 0: "jack", 2: "marry"}

	fmt.Println(numsArr1)
	fmt.Println(numsArr2)
	fmt.Println(names)

	for index, num := range numsArr1 {
		fmt.Println(index, num)
	}

	var sliceArr = []int{1, 2, 3}
	test(sliceArr)
	fmt.Println(sliceArr)
}

func test(a []int) {
	a[0] = 100
}
