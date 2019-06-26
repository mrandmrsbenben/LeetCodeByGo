package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	// preorder := []int{8, 5, 1, 7, 10, 12}
	// preorder := []int{1, 2, 3}
	preorder := []int{4, 5, 14, 20}
	fmt.Println("Output:", bstFromPreorder(preorder))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了95.65%的用户
//内存消耗 :3.3 MB, 在所有 Go 提交中击败了50.00%的用户
func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{preorder[0], nil, nil}
	switch len(preorder) {
	case 1:
		return root
	case 2:
		if preorder[1] < preorder[0] {
			root.Left = &TreeNode{preorder[1], nil, nil}
		} else {
			root.Right = &TreeNode{preorder[1], nil, nil}
		}
	case 3:
		if preorder[0] > preorder[1] && preorder[0] < preorder[2] {
			root.Left = &TreeNode{preorder[1], nil, nil}
			root.Right = &TreeNode{preorder[2], nil, nil}
		} else if preorder[0] < preorder[1] {
			root.Right = bstFromPreorder(preorder[1:])
		} else if preorder[0] > preorder[2] {
			root.Left = bstFromPreorder(preorder[1:])
		}
	default:
		if preorder[0] < preorder[1] {
			root.Right = bstFromPreorder(preorder[1:])
		} else if preorder[0] > preorder[len(preorder)-1] {
			root.Left = bstFromPreorder(preorder[1:])
		} else {
			for i := 1; i < len(preorder); i++ {
				if preorder[0] > preorder[i] &&
					i < len(preorder)-1 && preorder[0] < preorder[i+1] {
					root.Left = bstFromPreorder(preorder[1 : i+1])
					root.Right = bstFromPreorder(preorder[i+1:])
				}
			}
		}
	}
	return root
}
