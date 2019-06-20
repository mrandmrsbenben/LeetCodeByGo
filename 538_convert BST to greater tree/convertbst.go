package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
	"sort"
)

type TreeNode = common.TreeNode

func main() {
	root := common.MakeTree([]int{1, 1, 1})
	fmt.Printf("Output: %v\n", convertBST(root))
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	vals := make([]int, 0)
	var f1 func(r *TreeNode)
	f1 = func(r *TreeNode) {
		if r == nil {
			return
		}
		vals = append(vals, r.Val)
		f1(r.Left)
		f1(r.Right)
	}
	f1(root)
	sort.Ints(vals)
	sums := make(map[int]int)
	sums[vals[len(vals)-1]] = 0
	sum := 0
	for i := len(vals) - 2; i >= 0; i-- {
		if vals[i] < vals[i+1] {
			sum = sum + vals[i+1]
			sums[vals[i]] = sum
		}
	}
	fmt.Println(sums)
	var f2 func(r *TreeNode)
	f2 = func(r *TreeNode) {
		if r == nil {
			return
		}
		r.Val = r.Val + sums[r.Val]
		f2(r.Left)
		f2(r.Right)
	}
	f2(root)
	return root
}
