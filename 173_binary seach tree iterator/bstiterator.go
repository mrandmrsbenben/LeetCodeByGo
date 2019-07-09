package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

//执行用时 :64 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :10.4 MB, 在所有 Go 提交中击败了44.44%的用户
type BSTIterator struct {
	vals []int
}

func Constructor(root *TreeNode) BSTIterator {
	vals := make([]int, 0)
	var getVals func(node *TreeNode)
	getVals = func(node *TreeNode) {
		if node == nil {
			return
		}
		getVals(node.Left)
		vals = append(vals, node.Val)
		getVals(node.Right)
	}
	getVals(root)
	return BSTIterator{vals}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	val := this.vals[0]
	this.vals = this.vals[1:]
	return val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.vals) > 0
}

func main() {
	vals := []int{7, 3, 15, -1, -1, 9, 20}
	root := common.MakeTree(vals)
	obj := Constructor(root)
	fmt.Println("Output:", obj.Next())
	fmt.Println("Output:", obj.Next())
	fmt.Println("Output:", obj.HasNext())
	fmt.Println("Output:", obj.Next())
	fmt.Println("Output:", obj.HasNext())
	fmt.Println("Output:", obj.Next())
	fmt.Println("Output:", obj.HasNext())
	fmt.Println("Output:", obj.Next())
	fmt.Println("Output:", obj.HasNext())
}
