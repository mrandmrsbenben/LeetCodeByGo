package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// TreeNode TreeNode
type TreeNode = common.TreeNode

// https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
func main() {
	A := common.MakeTree([]int{3, 4, 5, 1, 2})
	B := common.MakeTree([]int{4, 1})
	fmt.Println("Output: ", isSubStructure(A, B))
}

//执行用时 :24 ms, 在所有 Go 提交中击败了91.49%的用户
//内存消耗 :6.6 MB, 在所有 Go 提交中击败了100.00%的用户
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if isSub(A, B) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSub(A, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	} else if A == nil {
		return false
	} else if B == nil {
		return true
	}
	if A.Val == B.Val {
		return isSub(A.Left, B.Left) && isSub(A.Right, B.Right)
	}
	return false
}
