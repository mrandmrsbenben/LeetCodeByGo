package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{1, 0, 1, 0, 1, 0, 1}
	root := common.MakeTree(vals)
	fmt.Printf("Output: %v\n", sumRootToLeaf(root))
}

//执行用时 :4 ms, 在所有Go提交中击败了96.55%
//的用户内存消耗 :3.4 MB, 在所有Go提交中击败了100.00%的用户
func sumRootToLeaf(root *TreeNode) int {
	sum := 0
	var fSum func(r *TreeNode, bits string)
	fSum = func(r *TreeNode, bits string) {
		if r == nil {
			return
		}
		if r.Val == 0 {
			bits = bits + "0"
		} else {
			bits = bits + "1"
		}
		if r.Left == nil && r.Right == nil {
			sum += binaryStrToInt(bits)
		} else {
			fSum(r.Left, bits)
			fSum(r.Right, bits)
		}
	}
	fSum(root, "")
	return sum
}

func binaryStrToInt(bs string) int {
	n := 0
	b := 1
	for i := len(bs) - 1; i >= 0; i-- {
		if bs[i] == '1' {
			n += b
		}
		b *= 2
	}
	return n
}
