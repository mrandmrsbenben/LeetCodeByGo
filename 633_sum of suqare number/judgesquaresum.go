package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Output:", judgeSquareSum(29))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了77.78%的用户
func judgeSquareSum(c int) bool {
	n := int(math.Sqrt(float64(c)))
	for i := 0; i <= n; i++ {
		if i*i+n*n == c {
			return true
		} else if i*i+n*n > c {
			n--
			continue
		}
	}
	return false
}
