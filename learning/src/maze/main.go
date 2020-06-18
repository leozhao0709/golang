package main

import (
	"fmt"
	"os"
)

func readMaze() [][]int8 {
	file, err := os.Open("src/maze/maze.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var row, column int8
	fmt.Fscanln(file, &row, &column)

	maze := make([][]int8, row)
	for i := range maze {
		maze[i] = make([]int8, column)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

func printMaze(maze [][]int8) {
	for i := range maze {
		for j := range maze[i] {
			fmt.Printf("%d ", maze[i][j])
		}
		fmt.Println()
	}
}

type point struct {
	i, j, val int8
	visited   bool
	parent    *point
}

func bfsWalk(maze [][]point) {
}

func main() {
	mazeData := readMaze()
	// printMaze(maze)

	maze := make([][]point, len(mazeData))
	for i := range maze {
		maze[i] = make([]point, len(maze[0]))
		for j := range maze[i] {
			maze[i][j] = point{i: int8(i), j: int8(j), val: mazeData[i][j]}
		}
	}

	bfsWalk(maze)
}
