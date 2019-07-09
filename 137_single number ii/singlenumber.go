package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := []int{2, 2, 3, 2}
	// nums := []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println("Output:", singleNumber(nums))
}

//执行用时 :8 ms, 在所有 Go 提交中击败了96.15%的用户
//内存消耗 :4.5 MB, 在所有 Go 提交中击败了8.33%的用户
func singleNumber(nums []int) int {
	nm := make(map[int]int)
	var single int
	for _, n := range nums {
		if _, ok := nm[n]; !ok {
			nm[n] = 1
		} else {
			nm[n]++
		}
	}
	for n, c := range nm {
		if c == 1 {
			single = n
			break
		}
	}
	return single
}

//执行用时 :504 ms, 在所有 Go 提交中击败了5.77%的用户
//内存消耗 :14.4 MB, 在所有 Go 提交中击败了8.33%的用户
func singleNumberSlow(nums []int) int {
	s := ""
	for _, n := range nums {
		s += "," + strconv.Itoa(n)
	}
	s += ","
	var substr string
	for _, n := range nums {
		substr = "," + strconv.Itoa(n) + ","
		if strings.Count(s, substr) == 1 {
			return n
		}
		// if strings.Index(s, substr) == strings.LastIndex(s, substr) {
		// 	return n
		// }
	}
	return 0
}
