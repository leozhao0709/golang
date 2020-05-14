package main

import (
	"fmt"
	"math"
)

func round(x float64, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func main() {
	var num1, num2 = 1.03, 0.42
	result := round(num1-num2, 0.01)
	fmt.Println(result)
}
