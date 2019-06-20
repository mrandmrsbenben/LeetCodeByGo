package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type ListNode = common.ListNode

func main() {
	str := ""
	fmt.Printf("Output: %v\n", deleteDuplicates(common.MakeListNode(str)))
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	} else {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
			deleteDuplicates(head)
		} else {
			deleteDuplicates(head.Next)
		}
	}
	return head
}
