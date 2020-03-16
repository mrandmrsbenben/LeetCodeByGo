package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// ListNode Linked List Node
type ListNode = common.ListNode

// https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
func main() {
	// l1 := common.MakeListNode("1->2->4")
	// l2 := common.MakeListNode("1->3->4")
	l1 := common.MakeListNode("-10->-9->-7->-1->-1->3->3->7->7")
	l2 := common.MakeListNode("-6->-4->1->1->2->6")
	fmt.Printf("Output: %v\n", mergeTwoLists(l1, l2))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var root, l *ListNode
	if l1.Val < l2.Val {
		root = l1
		l1 = l1.Next
	} else {
		root = l2
		l2 = l2.Next
	}
	l = root
	for l1 != nil || l2 != nil {
		if l1 == nil {
			l.Next = l2
			l2 = l2.Next
		} else if l2 == nil {
			l.Next = l1
			l1 = l1.Next
		} else if l1.Val < l2.Val {
			l.Next = l1
			l1 = l1.Next
		} else {
			l.Next = l2
			l2 = l2.Next
		}
		l = l.Next
	}
	return root
}

func mergeTwoListsV2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var l *ListNode
	if l1.Val < l2.Val {
		l = l1
		l.Next = mergeTwoLists(l1.Next, l2)
	} else {
		l = l2
		l.Next = mergeTwoLists(l1, l2.Next)
	}
	return l
}
