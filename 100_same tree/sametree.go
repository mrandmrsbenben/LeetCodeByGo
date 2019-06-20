package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
	"strconv"
)

type TreeNode = common.TreeNode

func main() {
	p := common.MakeTree([]int{1, 2, 1})
	q := common.MakeTree([]int{1, 1, 2})
	fmt.Printf("Output: %v\n", isSameTree(p, q))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSameTree0(p *TreeNode, q *TreeNode) bool {
	var f func(root *TreeNode, str string) string
	f = func(root *TreeNode, str string) string {
		if root == nil {
			return str
		}
		str = str + strconv.Itoa(root.Val) + ","
		if root.Left != nil || root.Right != nil {
			if root.Left != nil {
				str = f(root.Left, str)
			} else {
				str = str + "nil,"
			}
			if root.Right != nil {
				str = f(root.Right, str)
			} else {
				str = str + "nil,"
			}
		}
		return str
	}
	return f(p, "") == f(q, "")
}
