package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	date1 := "2019-01-09"
	fmt.Printf("Expect1: 9\n")
	fmt.Printf("Output1: %v\n", dayOfYear(date1))
	date2 := "2019-02-10"
	fmt.Printf("Expect2: 41\n")
	fmt.Printf("Output2: %v\n", dayOfYear(date2))
	date3 := "2003-03-01"
	fmt.Printf("Expect3: 60\n")
	fmt.Printf("Output3: %v\n", dayOfYear(date3))
	date4 := "2004-03-01"
	fmt.Printf("Expect4: 61\n")
	fmt.Printf("Output4: %v\n", dayOfYear(date4))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户
func dayOfYear(date string) int {
	md := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	s := strings.Split(date, "-")
	year, _ := strconv.Atoi(s[0])
	month, _ := strconv.Atoi(s[1])
	days, _ := strconv.Atoi(s[2])
	for i := 1; i < month; i++ {
		days += md[i]
	}
	if month > 2 && year%4 == 0 && year != 1900 {
		days++
	}
	return days
}
