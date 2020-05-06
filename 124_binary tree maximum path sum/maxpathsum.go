package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
	"math"
)

// TreeNode Binary Tree Node
type TreeNode = common.TreeNode

func main() {
	// root := common.MakeTree2([]int{1, 2, 3})
	// root := common.MakeTree2([]int{-10, 9, 20, -1, -1, 15, 7})
	root := common.MakeTree2([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1})
	fmt.Println("Output: ", maxPathSum(root))
}

//执行用时 :16 ms, 在所有 Go 提交中击败了97.93%的用户
//内存消耗 :6.7 MB, 在所有 Go 提交中击败了100.00%的用户
func maxPathSum(root *TreeNode) int {
	res := math.MinInt64

	var findMax func(node *TreeNode) int
	findMax = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := findMax(node.Left), findMax(node.Right)
		val := node.Val
		if l > 0 {
			val += l
		}
		if r > 0 {
			val += r
		}
		if res < val {
			res = val
		}
		return node.Val + max(l, r)
	}
	findMax(root)
	return res
}

func max(x, y int) int {
	if x < 0 && y < 0 {
		return 0
	}
	if x > y {
		return x
	}
	return y
}
