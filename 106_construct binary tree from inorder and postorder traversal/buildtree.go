package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// TreeNode Binary Tree Node
type TreeNode = common.TreeNode

func main() {
	// inorder := []int{9, 3, 15, 20, 7}
	// postorder := []int{9, 15, 7, 20, 3}
	// inorder := []int{1, 2, 3}
	// postorder := []int{3, 2, 1}
	// inorder := []int{1, 2, 3, 4}
	// postorder := []int{1, 4, 3, 2}
	inorder := []int{1, 2, 3, 4}
	postorder := []int{2, 1, 4, 3}
	fmt.Println("Output: ", buildTree(inorder, postorder))
}

//执行用时：4 ms, 在所有 Go 提交中击败了96.14%的用户
//内存消耗：4 MB, 在所有 Go 提交中击败了100.00%的用户
func buildTree(inorder []int, postorder []int) *TreeNode {
	if inorder == nil || len(inorder) == 0 || len(inorder) != len(postorder) {
		return nil
	} else if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0], Left: nil, Right: nil}
	}

	val := postorder[len(postorder)-1]
	i := 0
	for i = range inorder {
		if inorder[i] == val {
			break
		}
	}
	lio, rio := inorder[:i], inorder[i+1:]
	lpo, rpo := postorder[:len(lio)], postorder[len(lio):len(postorder)-1]
	return &TreeNode{Val: val, Left: buildTree(lio, lpo), Right: buildTree(rio, rpo)}
}
