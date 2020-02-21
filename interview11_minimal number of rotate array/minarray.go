package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{3, 4, 5, 1, 3}
	fmt.Println("Output: ", minArray(numbers))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了96.43%的用户
//内存消耗 :3.1 MB, 在所有 Go 提交中击败了100.00%的用户
func minArray(numbers []int) int {
	l := len(numbers)
	if l == 1 {
		return numbers[0]
	} else if l == 2 {
		if numbers[1] < numbers[0] {
			return numbers[1]
		}
		return numbers[0]
	} else if numbers[0] < numbers[l-1] {
		return numbers[0]
	}

	leftN := minArray(numbers[0 : l/2])
	rightN := minArray(numbers[l/2:])
	if leftN < rightN {
		return leftN
	}
	return rightN
}

//执行用时 :4 ms, 在所有 Go 提交中击败了96.43%的用户
//内存消耗 :3.1 MB, 在所有 Go 提交中击败了100.00%的用户
func minArrayV2(numbers []int) int {
	for i := len(numbers) - 1; i >= 1; i-- {
		if numbers[i] < numbers[i-1] {
			return numbers[i]
		}
	}
	return numbers[0]
}

//执行用时 :4 ms, 在所有 Go 提交中击败了96.43%的用户
//内存消耗 :3.1 MB, 在所有 Go 提交中击败了100.00%的用户
func minArrayV1(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}
