package main

import (
	"fmt"
	"os"
	"time"
)

func init() {
	os.Setenv("TZ", "Asia/Shanghai")
}

func main() {
	// basicTest()
	// timeZoneTest()

	fmt.Println(time.Now())
}

func timeZoneTest() {
	now := time.Now()
	fmt.Println(now)

	time.Now().UTC().Zone()
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("load location error", err)
		return
	}

	localTime := now.In(loc)
	fmt.Println(localTime)
}

func basicTest() {
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
