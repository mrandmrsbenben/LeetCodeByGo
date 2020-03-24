package main

import (
	"fmt"
	"math"
	"strconv"
)

// https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/
func main() {
	//0-9:1
	//10-99:10+9*1
	//100-999:100+9*20=280
	//1000-9999:1000+9*300=3700
	//10000-99999:10000+9*4000=46000
	fmt.Println("Output: ", countDigitOne(323))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了37.50%的用户
func countDigitOne(n int) int {
	if n > 0 {
		return count(strconv.Itoa(n))
	}
	return 0
}

func count(strN string) int {
	if len(strN) == 0 {
		return 0
	}
	first := int(strN[0] - '0')
	if len(strN) == 1 && first == 0 {
		return 0
	}
	if len(strN) == 1 && first > 0 {
		return 1
	}
	var numFirstDigit int
	if first > 1 {
		numFirstDigit = int(math.Pow10(len(strN) - 1))
	} else if first == 1 {
		post, _ := strconv.Atoi(strN[1:])
		numFirstDigit = post + 1
	}
	numOtherDigits := first * (len(strN) - 1) * int(math.Pow10(len(strN)-2))
	return numFirstDigit + numOtherDigits + count(strN[1:])
}
