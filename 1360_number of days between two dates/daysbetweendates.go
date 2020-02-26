package main

import (
	"fmt"
	"time"
)

func main() {
	date1 := "2019-06-29"
	date2 := "2019-06-30"
	fmt.Println("Output: ", daysBetweenDates(date1, date2))
}

func daysBetweenDates(date1 string, date2 string) int {
	d1, _ := time.Parse("2006-01-02", date1)
	d2, _ := time.Parse("2006-01-02", date2)
	days := (d2.Sub(d1)).Hours() / 24
	if days < 0 {
		days *= -1
	}
	return int(days)
}
