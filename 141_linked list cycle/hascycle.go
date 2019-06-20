package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type ListNode = common.ListNode

func main() {
	str := "3->2->0->-4"
	fmt.Printf("Output: %v\n", hasCycle(common.MakeListNode(str)))
}

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]int)
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = head.Val
		head = head.Next
	}
	return false
}
