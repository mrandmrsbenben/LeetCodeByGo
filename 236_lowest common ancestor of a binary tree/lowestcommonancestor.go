package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// TreeNode Binary Tree Node
type TreeNode = common.TreeNode

func main() {
	root := common.MakeTree2([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4})
	// p := common.MakeTree2([]int{5, 6, 2, -1, -1, 7, 4})
	// q := common.MakeTree2([]int{1, 0, 8})
	p := common.MakeTree2([]int{7})
	q := common.MakeTree2([]int{4})
	fmt.Println("Output: ", lowestCommonAncestor(root, p, q))
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var isLeaf func(r, n *TreeNode) bool
	isLeaf = func(r, n *TreeNode) bool {
		if r == nil {
			return false
		}

		if r.Val == n.Val {
			return true
		}

		return isLeaf(r.Left, n) || isLeaf(r.Right, n)
	}

	if isLeaf(p, q) {
		return p
	} else if isLeaf(q, p) {
		return q
	}

	var findLowest func(r *TreeNode) *TreeNode
	findLowest = func(r *TreeNode) *TreeNode {
		isPinL := isLeaf(r.Left, p)
		isQinR := isLeaf(r.Right, q)
		isPinR := isLeaf(r.Right, p)
		isQinL := isLeaf(r.Left, q)
		if (isPinL && isQinR) || (isPinR && isQinL) {
			return r
		} else if isPinL && isQinL {
			return findLowest(r.Left)
		} else {
			return findLowest(r.Right)
		}
	}
	return findLowest(root)
}
