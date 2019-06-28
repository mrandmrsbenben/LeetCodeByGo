package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type ListNode = common.ListNode
type TreeNode = common.TreeNode

func main() {
	str := "-10->-3"
	head := common.MakeListNode(str)
	fmt.Println("Output:", sortedListToBST(head))
}

//执行用时 :996 ms, 在所有 Go 提交中击败了68.12%的用户
//内存消耗 :35.1 MB, 在所有 Go 提交中击败了100.00%的用户
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	list := make([]int, 0)
	for {
		list = append(list, head.Val)
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	return makeBST(list)
}

func makeBST(vals []int) *TreeNode {
	var root *TreeNode
	switch len(vals) {
	case 0:
		root = nil
	case 2:
		root = &TreeNode{vals[1], &TreeNode{vals[0], nil, nil}, nil}
	case 3:
		root = &TreeNode{vals[1], &TreeNode{vals[0], nil, nil}, &TreeNode{vals[2], nil, nil}}
	default:
		mid := len(vals) / 2
		root = &TreeNode{vals[mid], makeBST(vals[0:mid]), makeBST(vals[mid+1:])}
	}
	return root
}
