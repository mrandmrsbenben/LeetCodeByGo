package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// TreeNode Binary Tree Node
type TreeNode = common.TreeNode

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	fmt.Println("Output:", buildTree(preorder, inorder))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :4 MB, 在所有 Go 提交中击败了100.00%的用户
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0], Left: nil, Right: nil}
	if len(preorder) == 1 {
		return root
	}
	i := 0
	for i = range inorder {
		if inorder[i] == root.Val {
			break
		}
	}
	leftin, rightin := inorder[0:i], inorder[i+1:]
	leftpre, rightpre := preorder[1:1+i], preorder[1+i:]
	root.Left = buildTree(leftpre, leftin)
	root.Right = buildTree(rightpre, rightin)
	return root
}
