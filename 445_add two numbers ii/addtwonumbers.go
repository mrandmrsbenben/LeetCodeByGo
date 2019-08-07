package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type ListNode = common.ListNode

func main() {
	// l1 := common.MakeListNode("7->2->4->3")
	// l2 := common.MakeListNode("5->6->4")
	l1 := common.MakeListNode("1->NULL")
	l2 := common.MakeListNode("9->9->NULL")
	fmt.Println("Output:", addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	n1 := listToArray(l1)
	n2 := listToArray(l2)
	if len(n2) > len(n1) {
		n1, n2 = n2, n1
	}
	add := 0
	n := make([]int, len(n1))
	for i, j := len(n1)-1, len(n2)-1; i >= 0; i-- {
		if j >= 0 {
			n[i] = (n1[i] + n2[j] + add) % 10
			add = (n1[i] + n2[j] + add) / 10
			j--
		} else {
			n[i], add = (n1[i]+add)%10, (n1[i]+add)/10
		}
	}
	if add > 0 {
		n = append([]int{add}, n...)
	}
	var l, head *ListNode
	for i := range n {
		if l == nil {
			head = &ListNode{n[i], nil}
			l = head
		} else {
			l.Next = &ListNode{n[i], nil}
			l = l.Next
		}
	}
	return head
}

func listToArray(l *ListNode) []int {
	n := make([]int, 0)
	for l != nil {
		n = append(n, l.Val)
		l = l.Next
	}
	return n
}
