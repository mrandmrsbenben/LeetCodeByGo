// https://leetcode-cn.com/problems/remove-duplicate-node-lcci/
package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// ListNode Node of Linked List
type ListNode = common.ListNode

func main() {
	vals := []int{1, 1, 1, 1, 2}
	root := common.MakeListNodeByArray(vals)
	fmt.Println("Output: ", removeDuplicateNodes(root))
}

//执行用时 :12 ms, 在所有 Go 提交中击败了96.17%的用户
//内存消耗 :6.1 MB, 在所有 Go 提交中击败了100.00%的用户
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	vmap := make(map[int]int)
	vmap[head.Val] = 1
	node := head
	for node.Next != nil {
		if vmap[node.Next.Val] == 0 {
			vmap[node.Next.Val] = 1
			node = node.Next
		} else {
			if node.Next.Next == nil {
				node.Next = nil
			} else {
				node.Next.Val = node.Next.Next.Val
				node.Next.Next = node.Next.Next.Next
			}
		}
	}
	return head
}
