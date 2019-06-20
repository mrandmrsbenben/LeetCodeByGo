package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type ListNode = common.ListNode

func main() {
	input := "1->2->3->4->5->NULL"
	fmt.Printf("Output: %v\n", reverseList(common.MakeListNode(input)))
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var root, prev *ListNode
	prev = &ListNode{head.Val, nil}
	for head.Next != nil {
		// vals = append(vals, head.Next.Val)
		// head = head.Next
		root = &ListNode{head.Next.Val, prev}
		prev = root
		head = head.Next
	}
	return root
}
