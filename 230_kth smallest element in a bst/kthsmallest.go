package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	k := 2
	vals := []int{3, 1, 4, -1, 2}
	// k := 3
	// vals := []int{5, 3, 6, 2, 4, -1, -1, 1}
	root := common.MakeTree(vals)
	fmt.Println("Output:", kthSmallest(root, k))
}

//执行用时 :12 ms, 在所有 Go 提交中击败了99.30%的用户
//内存消耗 :6 MB, 在所有 Go 提交中击败了56.10%的用户
func kthSmallest(root *TreeNode, k int) int {
	kv := 0
	var findKth func(r *TreeNode)
	findKth = func(r *TreeNode) {
		if r == nil || k == 0 {
			return
		}
		findKth(r.Left)
		k--
		if k == 0 {
			kv = r.Val
			return
		}
		findKth(r.Right)
	}
	findKth(root)
	return kv
}
