package main

import "fmt"

func main() {
	nums := []int{1, 1, 3, 4, 5}
	fmt.Println("Output: ", findMagicIndex(nums))
}

//执行用时：12 ms, 在所有 Go 提交中击败了96.88%的用户
//内存消耗：5.9 MB, 在所有 Go 提交中击败了100.00%的用户
func findMagicIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	var find func([]int, int) int
	find = func(n []int, start int) int {
		if len(n) == 0 {
			return -1
		}

		i := (len(n) - 1) / 2
		if n[i] == i+start {
			return i + start
		}

		l := find(n[:i], start)
		if l != -1 {
			return l
		}
		return find(n[i+1:], start+i+1)
	}

	return find(nums, 0)
}
