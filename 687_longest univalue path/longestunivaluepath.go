package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{5, 5, 5, 1, 1, -1, 5}
	fmt.Println("Output:", longestUnivaluePath(common.MakeTree(vals)))
}

//执行用时 :120 ms, 在所有Go提交中击败了86.67%的用户
//内存消耗 :16.3 MB, 在所有Go提交中击败了5.56%的用户
func longestUnivaluePath(root *TreeNode) int {
	max := 0
	var findMax func(r *TreeNode)
	findMax = func(r *TreeNode) {
		if r == nil {
			return
		}
		path := findVal(r.Left, r.Val) + findVal(r.Right, r.Val)
		if path > max {
			max = path
		}
		findMax(r.Left)
		findMax(r.Right)
	}
	findMax(root)
	return max
}

func findVal(r *TreeNode, val int) int {
	if r == nil || r.Val != val {
		return 0
	}
	left := findVal(r.Left, val)
	right := findVal(r.Right, val)
	if right > left {
		return right + 1
	}
	return left + 1
}
