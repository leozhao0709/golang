package main

import (
	"fmt"
	"math"
)

func round(x float64, unit float64) float64 {
	unit = 1 / unit
	return math.Round(x*unit) / unit
}

func main() {
	var num1, num2 = 1.03, 0.42
	result1 := round(num1-num2, 0.1)
	fmt.Println(result1)
}
