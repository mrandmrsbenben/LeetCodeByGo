package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	// s := common.MakeTree([]int{3, 4, 5, 1, -1, 2})
	// t := common.MakeTree([]int{3, 1, 2})
	// s := common.MakeTree([]int{1, 1})
	// t := common.MakeTree([]int{1})
	s := common.MakeTree([]int{3, 4, 5, 1, 2})
	t := common.MakeTree([]int{3, -1, 5})
	fmt.Printf("Output: %v\n", isSubtree(s, t))
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	} else if s == nil || t == nil {
		return false
	}
	if s.Val == t.Val {
		if isSub(s.Left, t.Left) && isSub(s.Right, t.Right) {
			return true
		}
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSub(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	} else if s == nil || t == nil {
		return false
	}
	if s.Val == t.Val {
		if isSubtree(s.Left, t.Left) && isSubtree(s.Right, t.Right) {
			return true
		}
	}
	return false
}
