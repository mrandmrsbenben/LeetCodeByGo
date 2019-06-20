package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
	"strconv"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{1, 2, 3, -1, 5}
	fmt.Printf("Output :%v\n", binaryTreePaths(common.MakeTree(vals)))
}

func binaryTreePaths(root *TreeNode) []string {
	paths := make([]string, 0)
	var f func(t *TreeNode, str string)
	f = func(t *TreeNode, str string) {
		if t == nil {
			return
		}
		str = str + "->" + strconv.Itoa(t.Val)
		if t.Left == nil && t.Right == nil {
			paths = append(paths, str[2:])
			return
		}
		if t.Left != nil {
			f(t.Left, str)
		}
		if t.Right != nil {
			f(t.Right, str)
		}
	}
	f(root, "")
	return paths
}
