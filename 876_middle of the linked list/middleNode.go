package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type ListNode = common.ListNode

func main() {
	head := common.MakeListNodeByArray([]int{1, 2, 3, 4, 5})
	fmt.Printf("Output: %v\n", middleNode(head))
}

func middleNode(head *ListNode) *ListNode {
	nodes := make([]*ListNode, 0)
	for {
		nodes = append(nodes, head)
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	return nodes[len(nodes)/2]
}
