package main

import "fmt"

func main() {
	target := 7
	fmt.Printf("Output: %d\n", reachNumber(target))
}

// 执行用时 : 0 ms, 在Reach a Number的Go提交中击败了100.00% 的用户
// 内存消耗 : 2 MB, 在Reach a Number的Go提交中击败了50.00% 的用户
func reachNumber(target int) int {
	if target < 0 {
		target *= -1
	}
	if target == 1 {
		return 1
	} else if target == 2 {
		return 3
	}
	sum := 0
	steps := 0
	for {
		steps++
		sum += steps
		if sum < target {
			continue
		}
		if sum == target || (sum-target)%2 == 0 {
			return steps
		} else if steps%2 == 1 {
			steps += 2
			return steps
		} else {
			steps++
			return steps
		}
	}
}
