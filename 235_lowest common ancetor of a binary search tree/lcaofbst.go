package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	// r := []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}
	// p := common.MakeTree([]int{2, 0, 4, -1, -1, 3, 5})
	// // q := common.MakeTree([]int{4, 3, 5})
	// q := common.MakeTree([]int{8, 7, 9})
	r := []int{2, 1, 3}
	p := common.MakeTree([]int{2})
	q := common.MakeTree([]int{3})
	root := common.MakeTree(r)
	fmt.Printf("Output: %v\n", lowestCommonAncestor(root, p, q).Val)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//find p
	pParents := make([]*TreeNode, 0)
	var find func(root, target *TreeNode, parents []*TreeNode) []*TreeNode
	find = func(root, target *TreeNode, parents []*TreeNode) []*TreeNode {
		if root == nil {
			return nil
		}
		if root.Val == target.Val {
			parents = append(parents, root)
			return parents
		}
		left := find(root.Left, target, parents)
		if left != nil {
			parents = append(left, root)
			return parents
		}
		right := find(root.Right, target, parents)
		if right != nil {
			parents = append(right, root)
			return parents
		}
		return nil
	}
	pParents = find(root, p, pParents)
	var findq func(root, target *TreeNode) *TreeNode
	findq = func(root, target *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if root.Val == target.Val {
			return root
		}
		left := findq(root.Left, target)
		if left != nil {
			return left
		}
		return findq(root.Right, target)
	}
	for _, node := range pParents {
		if findq(node, q) != nil {
			return node
		}
	}
	return root
}
