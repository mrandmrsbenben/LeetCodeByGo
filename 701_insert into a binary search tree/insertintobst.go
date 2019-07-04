package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	val := 5
	vals := []int{4, 2, 7, 1, 3}
	root := common.MakeTree(vals)
	fmt.Println("Output:", insertIntoBST(root, val))
}

//执行用时 :860 ms, 在所有 Go 提交中击败了87.10%的用户
//内存消耗 :91 MB, 在所有 Go 提交中击败了16.67%的用户
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	r := root
	for {
		if r.Val > val {
			if r.Left == nil {
				r.Left = &TreeNode{val, nil, nil}
				break
			} else {
				r = r.Left
			}
		} else {
			if r.Right == nil {
				r.Right = &TreeNode{val, nil, nil}
				break
			} else {
				r = r.Right
			}
		}
	}
	return root
}

//执行用时 :928 ms, 在所有 Go 提交中击败了61.29%的用户
//内存消耗 :143.2 MB, 在所有 Go 提交中击败了16.67%的用户
func insertIntoBSTV1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > val {
		if root.Left != nil {
			insertIntoBST(root.Left, val)
		} else {
			root.Left = &TreeNode{val, nil, nil}
		}
	} else {
		if root.Right != nil {
			insertIntoBST(root.Right, val)
		} else {
			root.Right = &TreeNode{val, nil, nil}
		}
	}
	return root
}
