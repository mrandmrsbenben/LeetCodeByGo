package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type ListNode = common.ListNode

func main() {
	l1 := common.MakeListNode("2->4->3")
	l2 := common.MakeListNode("5->6->4")
	fmt.Println("Output:", addTwoNumbers(l1, l2))
}

//执行用时 :16 ms, 在所有 Go 提交中击败了76.96%的用户
//内存消耗 :5.2 MB, 在所有 Go 提交中击败了25.83%的用户
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var l, head *ListNode
	val, add := 0, 0
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			val = l1.Val + l2.Val + add
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 == nil {
			val = l2.Val + add
			l2 = l2.Next
		} else {
			val = l1.Val + add
			l1 = l1.Next
		}
		add = val / 10
		if l == nil {
			head = &ListNode{val % 10, nil}
			l = head
		} else {
			l.Next = &ListNode{val % 10, nil}
			l = l.Next
		}
	}
	if add > 0 {
		l.Next = &ListNode{add, nil}
	}
	return head
}
