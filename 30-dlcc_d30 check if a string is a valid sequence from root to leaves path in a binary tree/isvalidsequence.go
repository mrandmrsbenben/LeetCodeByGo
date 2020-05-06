package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

//TreeNode node of binary tree
type TreeNode = common.TreeNode

func main() {
	// root := common.MakeTree2([]int{0, 1, 0, 0, 1, 0, -1, -1, 1, 0, 0})
	// arr := []int{0, 1, 0, 1}
	// root := common.MakeTree2([]int{0, 1, 0, 0, 1, 0, -1, -1, 1, 0, 0})
	// arr := []int{0, 0, 1}
	root := common.MakeTree2([]int{0, 1, 0, 0, 1, 0, -1, -1, 1, 0, 0})
	arr := []int{0, 1, 1}
	fmt.Println("Output: ", isValidSequence(root, arr))
}

func isValidSequence(root *TreeNode, arr []int) bool {
	if root == nil {
		return len(arr) == 0
	}

	if root.Val != arr[0] {
		return false
	} else if len(arr) == 1 {
		return root.Left == nil && root.Right == nil
	}
	return isValidSequence(root.Left, arr[1:]) || isValidSequence(root.Right, arr[1:])
}
