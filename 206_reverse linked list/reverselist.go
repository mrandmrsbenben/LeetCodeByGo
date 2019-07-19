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
	node := reverseList(head.Next)
	head.Next.Next, head.Next = head, nil
	return node
}

func reverseListV1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, cur, next *ListNode
	cur = head
	for cur != nil {
		next = cur.Next
		cur.Next, prev = prev, cur
		cur = next
	}
	return prev
}
