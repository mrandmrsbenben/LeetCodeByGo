package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// TreeNode Binary Tree Node
type TreeNode = common.TreeNode

func main() {
	vals := []int{1, 2, 2, 3, 4, 4, 3, 5}
	fmt.Println("Output: ", isSymmetric(common.MakeTree(vals)))
}

//执行用时 :56 ms, 在所有 Go 提交中击败了78.67%的用户
//内存消耗 :19.2 MB, 在所有 Go 提交中击败了100.00%的用户
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	} else if root.Left == nil && root.Right == nil {
		return true
	} else if root.Left == nil || root.Right == nil {
		return false
	}

	nodes := []*TreeNode{root.Left, root.Right}
	hasLeaf := true
	var leaves []*TreeNode
	var i, j int
	for hasLeaf {
		hasLeaf = false
		leaves = make([]*TreeNode, len(nodes)*2)
		for i = 0; i < len(nodes)/2; i++ {
			j = len(nodes) - 1 - i
			if !equals(nodes[i], nodes[j]) {
				return false
			}
			if nodes[i] != nil {
				leaves[i*2], leaves[i*2+1] = nodes[i].Left, nodes[i].Right
				hasLeaf = true
			}
			if nodes[j] != nil {
				leaves[j*2], leaves[j*2+1] = nodes[j].Left, nodes[j].Right
				hasLeaf = true
			}
		}
		nodes = leaves
	}
	return true
}

func equals(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if l == nil || r == nil {
		return false
	}
	return l.Val == r.Val
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.9 MB, 在所有 Go 提交中击败了100.00%的用户
func isSymmetricV1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if l == nil || r == nil {
		return false
	}

	if l.Val == r.Val {
		return isMirror(l.Left, r.Right) && isMirror(l.Right, r.Left)
	}
	return false
}
