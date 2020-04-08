package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// TreeNode Binary Tree Node
type TreeNode = common.TreeNode

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
func main() {
	// vals := []int{3, 1, 4, -1, 2}
	vals := []int{5, 3, 6, 2, 4, -1, -1, 1}
	fmt.Println("Output: ", kthLargest(common.MakeTree(vals), 4))
}

//执行用时 :12 ms, 在所有 Go 提交中击败了86.00%的用户
//内存消耗 :6 MB, 在所有 Go 提交中击败了100.00%的用户
func kthLargest(root *TreeNode, k int) int {
	res := 0
	cnt := 0
	var findKth func(r *TreeNode)
	findKth = func(r *TreeNode) {
		if r == nil || cnt >= k {
			return
		}
		findKth(r.Right)
		cnt++
		if cnt == k {
			res = r.Val
			return
		}
		findKth(r.Left)
	}
	findKth(root)
	return res
}
