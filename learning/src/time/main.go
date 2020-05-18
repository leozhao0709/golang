package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	year := now.Year()
	month := now.Month()
	day := now.Day()

	fmt.Println("year: ", year, "month: ", int(month), "day: ", day)

	dateFormat := now.Format("2006/01/02 15:04:05.000")
	fmt.Println(dateFormat)

	t1 := time.Now().AddDate(0, 0, 1)
	duration := t1.Sub(now)
	fmt.Println("duration hours is", duration.Seconds())
}
