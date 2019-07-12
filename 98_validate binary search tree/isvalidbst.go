package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
	"math"
)

type TreeNode = common.TreeNode

func main() {
	// vals := []int{2, 1, 3}
	vals := []int{3, 1, 5, -1, -1, 4, 6}
	fmt.Println("Output:", isValidBST(common.MakeTree(vals)))
}

//执行用时 :8 ms, 在所有 Go 提交中击败了97.39%的用户
//内存消耗 :5.4 MB, 在所有 Go 提交中击败了80.58%的用户
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	prev := math.MinInt64
	var isValid func(node *TreeNode) bool
	isValid = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !isValid(node.Left) {
			return false
		}
		if node.Val > prev {
			prev = node.Val
		} else {
			return false
		}
		return isValid(node.Right)
	}
	return isValid(root)
}
