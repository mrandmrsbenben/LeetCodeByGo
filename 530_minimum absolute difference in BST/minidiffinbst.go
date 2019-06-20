package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
	"sort"
)

type TreeNode = common.TreeNode

func main() {
	root := common.MakeTree([]int{1, 3, 2})
	fmt.Printf("Output: %d\n", getMinimumDifference(root))
}

func getMinimumDifference(root *TreeNode) int {
	vals := make([]int, 0)
	var f func(r *TreeNode)
	f = func(r *TreeNode) {
		if r != nil {
			vals = append(vals, r.Val)
			f(r.Left)
			f(r.Right)
		}
	}
	f(root)
	sort.Ints(vals)
	min := vals[len(vals)-1]
	for i := 0; i < len(vals)-1; i++ {
		if min > vals[i+1]-vals[i] {
			min = vals[i+1] - vals[i]
		}
	}
	return min
}
