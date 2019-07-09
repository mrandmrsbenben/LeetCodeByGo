package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	// vals := []int{3, 0, 0}
	// vals := []int{0, 3, 0}
	// vals := []int{1, 0, 2}
	vals := []int{1, 0, 0, -1, 3}
	fmt.Println("Output:", distributeCoins(common.MakeTree(vals)))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :3.4 MB, 在所有 Go 提交中击败了33.33%的用户
func distributeCoins(root *TreeNode) int {
	if root == nil {
		return 0
	}
	moves := 0
	var distribute func(r *TreeNode) int
	distribute = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		left, right := distribute(r.Left), distribute(r.Right)
		moves += abs(left) + abs(right)
		return left + right + r.Val - 1
	}
	distribute(root)
	return moves
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}
