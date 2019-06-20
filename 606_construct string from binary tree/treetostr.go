package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
	"strconv"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{1, -1, 3}
	t := common.MakeTree(vals)
	fmt.Printf("Output: %s\n", tree2str(t))
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	str := ""
	var f func(root *TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		str = str + "(" + strconv.Itoa(root.Val)
		if root.Left != nil {
			f(root.Left)
		}
		if root.Right != nil {
			if root.Left == nil {
				str = str + "()"
			}
			f(root.Right)
		}
		str = str + ")"
	}
	f(t)
	return str[1 : len(str)-1]
}
