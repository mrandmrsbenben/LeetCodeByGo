package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type ListNode = common.ListNode

func main() {
	str := "1->2->3->4->5"
	n := 2
	fmt.Println("Output:", removeNthFromEnd(common.MakeListNode(str), n))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.3 MB, 在所有 Go 提交中击败了20.73%的用户
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := make([]*ListNode, 0)
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	if n == 1 {
		if len(nodes) == 1 {
			return nil
		}
		nodes[len(nodes)-2].Next = nil
	} else {
		deleteNode := nodes[len(nodes)-n]
		deleteNode.Val = nodes[len(nodes)-n+1].Val
		deleteNode.Next = nodes[len(nodes)-n+1].Next
	}
	return nodes[0]
}
