package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
	"sort"
)

type ListNode = common.ListNode

func main() {
	l1 := common.MakeListNode("1->2->4")
	l2 := common.MakeListNode("1->3->4")
	fmt.Printf("Output: %v\n", mergeTwoLists(l1, l2))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		if l2 == nil {
			return nil
		}
		return l2
	} else if l2 == nil {
		return l1
	}
	var f func(l *ListNode, vals []int) []int
	f = func(l *ListNode, vals []int) []int {
		vals = append(vals, l.Val)
		if l.Next != nil {
			vals = f(l.Next, vals)
		}
		return vals
	}
	vals := append(f(l1, make([]int, 0)), f(l2, make([]int, 0))...)
	sort.Ints(vals)
	nodes := make([]*ListNode, len(vals))
	for i := range vals {
		nodes[i] = &ListNode{vals[i], nil}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}
	return nodes[0]
}
