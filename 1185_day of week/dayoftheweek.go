package main

import "fmt"

func main() {
	// fmt.Println("Output: ", dayOfTheWeek(31, 8, 2019))
	// fmt.Println("Output: ", dayOfTheWeek(14, 2, 1971))
	// fmt.Println("Output: ", dayOfTheWeek(18, 7, 1999))
	// fmt.Println("Output: ", dayOfTheWeek(15, 8, 1993))
	// fmt.Println("Output: ", dayOfTheWeek(16, 1, 1979))
	fmt.Println("Output: ", dayOfTheWeek(29, 2, 2016))
}

func dayOfTheWeek(day int, month int, year int) string {
	weekdays := []string{"Thursday", "Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday"}
	monthdays := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	days := day - 1

	// Calculate the days of months
	for i := 1; i < month; i++ {
		days += monthdays[i]
	}
	if month > 2 && isSpecialYear(year) {
		days++
	}

	// Calculate the days of years
	for i := 1970; i < year; i++ {
		if isSpecialYear(i) {
			days += 366
		} else {
			days += 365
		}
	}
	return weekdays[days%7]
}

func isSpecialYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}
