package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// input := "2-1-1"
	input := "2*3-4*5"
	fmt.Println("Output:", diffWaysToCompute(input))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :6.5 MB, 在所有 Go 提交中击败了66.67%的用户
func diffWaysToCompute(input string) []int {
	// convert string to numbers
	numbers, opers := transToArray(input)
	switch len(numbers) {
	case 0:
		return []int{}
	case 1:
		return []int{numbers[0]}
	case 2:
		return []int{caculate(numbers[0], numbers[1], opers[0])}
	}
	// compute all the different result
	var compute func(nums []int, ops []string) []int
	compute = func(nums []int, ops []string) []int {
		if len(nums) == 1 {
			return nums
		} else if len(nums) == 2 {
			return []int{caculate(nums[0], nums[1], ops[0])}
		}
		result := make([]int, 0)
		for i := range ops {
			l := compute(nums[0:i+1], ops[0:i])
			r := compute(nums[i+1:], ops[i+1:])
			rs := make([]int, len(l)*len(r))
			index := 0
			for _, m := range l {
				for _, n := range r {
					rs[index] = caculate(m, n, ops[i])
					index++
				}
			}
			result = append(result, rs...)
		}
		return result
	}
	return compute(numbers, opers)
}

// caculator
func caculate(m, n int, oper string) int {
	switch oper {
	case "+":
		return m + n
	case "-":
		return m - n
	case "*":
		return m * n
	}
	return 0
}

// convert string to numbers and operators
func transToArray(input string) ([]int, []string) {
	opers := make([]string, 0)
	sep := ","
	buf := strings.Replace(input, "+", sep, -1)
	buf = strings.Replace(buf, "-", sep, -1)
	buf = strings.Replace(buf, "*", sep, -1)
	ns := strings.Split(buf, sep)
	numbers := make([]int, len(ns))
	l := 0
	for i, n := range ns {
		numbers[i], _ = strconv.Atoi(n)
		if i < len(ns)-1 {
			l += len(n) + 1
			opers = append(opers, input[l-1:l])
		}
	}
	return numbers, opers
}
