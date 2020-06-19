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
	i, j int
}

func reverse(slice []point) []point {
	i := 0
	j := len(slice) - 1
	for i <= j {
		slice[i], slice[j] = slice[j], slice[i]
		i++
		j--
	}
	return slice
}

func bfsWalk(mazeData [][]int8, start, end point) []point {
	rows := len(mazeData)
	columns := len(mazeData[0])
	steps := make([][]int, rows)
	for i := range steps {
		steps[i] = make([]int, columns)
	}
	steps[0][0] = 1

	queue := []point{start}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == end {
			break
		}

		up := curr.i - 1
		down := curr.i + 1
		left := curr.j - 1
		right := curr.j + 1

		if up >= 0 && mazeData[up][curr.j] != 1 && steps[up][curr.j] == 0 {
			queue = append(queue, point{up, curr.j})
			steps[up][curr.j] = steps[curr.i][curr.j] + 1
		}

		if down < rows && mazeData[down][curr.j] != 1 && steps[down][curr.j] == 0 {
			queue = append(queue, point{down, curr.j})
			steps[down][curr.j] = steps[curr.i][curr.j] + 1
		}

		if left >= 0 && mazeData[curr.i][left] != 1 && steps[curr.i][left] == 0 {
			queue = append(queue, point{curr.i, left})
			steps[curr.i][left] = steps[curr.i][curr.j] + 1
		}

		if right < columns && mazeData[curr.i][right] != 1 && steps[curr.i][right] == 0 {
			queue = append(queue, point{curr.i, right})
			steps[curr.i][right] = steps[curr.i][curr.j] + 1
		}
	}

	for i := range steps {
		for j := range steps[i] {
			fmt.Printf("%3d", steps[i][j])
		}
		fmt.Println()
	}

	if steps[rows-1][columns-1] == 0 {
		return nil
	}

	shortestPaths := []point{end}
	step := steps[rows-1][columns-1]
	p := end
	for step > 1 {
		up := p.i - 1
		down := p.i + 1
		left := p.j - 1
		right := p.j + 1
		if up >= 0 && steps[up][p.j] == step-1 {
			p = point{up, p.j}
		} else if down < rows && steps[down][p.j] == step-1 {
			p = point{down, p.j}
		} else if left >= 0 && steps[p.i][left] == step-1 {
			p = point{p.i, left}
		} else if right < columns && steps[p.i][right] == step-1 {
			p = point{p.i, right}
		}
		shortestPaths = append(shortestPaths, p)
		step = step - 1
	}
	return reverse(shortestPaths)
}

func main() {
	mazeData := readMaze()
	// printMaze(mazeData)
	shortestPath := bfsWalk(mazeData, point{0, 0}, point{len(mazeData) - 1, len(mazeData[0]) - 1})
	fmt.Println(shortestPath)
}
