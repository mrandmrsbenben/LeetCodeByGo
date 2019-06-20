package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Printf("Output:%v\n", sortedArrayToBST(nums))
}

func sortedArrayToBST(nums []int) *TreeNode {
	var root *TreeNode
	switch len(nums) {
	case 0:
		return nil
	default:
		md := len(nums) / 2
		root = &TreeNode{nums[md], sortedArrayToBST(nums[0:md]), sortedArrayToBST(nums[md+1:])}
	}
	return root
}
