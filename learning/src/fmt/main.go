package main

import (
	"fmt"
	"os"
)

func main() {
	// scanfTest()
	// scanlnTest()
	fscanfTest()
}

func scanlnTest() {
	var s1 string
	var s2 string
	fmt.Scanln(&s1, &s2)
	fmt.Printf("s1 is %v, s2 is %v\n", s1, s2)
}

func scanfTest() {
	var s1 string
	var s2 string
	fmt.Scanf("%v", &s1)
	fmt.Scanf("%v", &s2)
	fmt.Printf("s1 is %v, s2 is %v\n", s1, s2)
}

func fscanfTest() {
	file, err := os.Open("./src/fmt/test.txt")
	if err != nil {
		fmt.Println("..open file fail", err)
		return
	}
	defer file.Close()

	var maze = [4][4]int{}

	for i := range maze {
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	fmt.Println(maze)
}
