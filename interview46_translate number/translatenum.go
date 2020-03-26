package main

import (
	"fmt"
	"strconv"
)

// https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
func main() {
	fmt.Println("Output: ", translateNum(506))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func translateNum(num int) int {
	mr := make(map[int]int)
	str := strconv.Itoa(num)
	l := len(str)

	var translate func(i int) int
	translate = func(i int) int {
		if n, ok := mr[i]; ok {
			return n
		}

		var res int
		switch l - i {
		case 1:
			res = 1
		case 2:
			if str[i:i+1] > "0" && str[i:] <= "25" {
				res = 2
			} else {
				res = 1
			}
		default:
			res = translate(i + 1)
			if str[i:i+1] > "0" && str[i:i+2] <= "25" {
				res += translate(i + 2)
			}
		}
		mr[i] = res
		return res
	}
	return translate(0)
}
