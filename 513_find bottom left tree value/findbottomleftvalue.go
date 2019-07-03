package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	// vals := []int{2, 1, 3}
	// vals := []int{2, 1, 3, 4, -1, 5, 6, -1, -1, -1, -1, 7}
	vals := []int{2, -1, 3}
	fmt.Println("Output:", findBottomLeftValue(common.MakeTree(vals)))
}

//执行用时 :12 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :5.4 MB, 在所有 Go 提交中击败了100.00%的用户
func findBottomLeftValue(root *TreeNode) int {
	var blv int
	bd := 0
	var getLeftValue func(r *TreeNode, depth int)
	getLeftValue = func(r *TreeNode, depth int) {
		if r == nil {
			return
		}
		getLeftValue(r.Left, depth+1)
		if depth > bd {
			bd = depth
			blv = r.Val
		}
		getLeftValue(r.Right, depth+1)
	}
	getLeftValue(root, 1)
	return blv
}
