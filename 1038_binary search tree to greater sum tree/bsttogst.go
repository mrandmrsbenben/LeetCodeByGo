package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8}
	fmt.Println("Output:", bstToGst(common.MakeTree(vals)))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.7 MB, 在所有 Go 提交中击败了100.00%的用户
func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	sum := 0
	var b2g func(r *TreeNode)
	b2g = func(r *TreeNode) {
		if r.Right != nil {
			b2g(r.Right)
		}
		r.Val += sum
		sum = r.Val
		if r.Left != nil {
			b2g(r.Left)
		}
		if r.Left == nil && r.Right == nil && sum == 0 {
			sum = r.Val
		}
	}
	b2g(root)
	return root
}
