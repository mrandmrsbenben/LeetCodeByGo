package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Output:", countBits(31))
}

//执行用时 :2108 ms, 在所有 Go 提交中击败了93.94%的用户
//内存消耗 :35.4 MB, 在所有 Go 提交中击败了6.06%的用户
func countBits(num int) []int {
	bits := make([]int, num+1)
	bits[0] = 0
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			bits[i] = bits[i/2]
		} else {
			bits[i] = bits[i-1] + 1
		}
	}
	return bits
}

//执行用时 :2148 ms, 在所有 Go 提交中击败了92.93%的用户
//内存消耗 :33.7 MB, 在所有 Go 提交中击败了6.06%的用户
func countBitsV1(num int) []int {
	bits := make([]int, num+1)
	for i := 0; i <= num; i++ {
		bits[i] = strings.Count(strconv.FormatInt(int64(i), 2), "1")
	}
	return bits
}
