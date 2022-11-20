package main

import (
	"fmt"
	"time"
)

func main() {
	var currentTime time.Time = time.Now()
	var year int = currentTime.Year()
	var month time.Month = currentTime.Month()
	var day int = currentTime.Day()
	var weekDay time.Weekday = currentTime.Weekday()
	fmt.Println(year, month, day, weekDay)
}
