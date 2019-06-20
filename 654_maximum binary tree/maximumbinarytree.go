package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	fmt.Println("Output:", constructMaximumBinaryTree(nums))
}

//执行用时 :108 ms, 在所有 Go 提交中击败了97.56%的用户
//内存消耗 :8.6 MB, 在所有 Go 提交中击败了81.82%的用户
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var root *TreeNode
	maxI := findIndexOfMaxValue(nums)
	root = &TreeNode{nums[maxI], nil, nil}
	left := nums[0:maxI]
	right := nums[maxI+1:]
	if len(left) > 0 {
		root.Left = constructMaximumBinaryTree(left)
	}
	if len(right) > 0 {
		root.Right = constructMaximumBinaryTree(right)
	}
	return root
}

func findIndexOfMaxValue(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	maxV, maxI := nums[0], 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxV {
			maxV = nums[i]
			maxI = i
		}
	}
	return maxI
}
