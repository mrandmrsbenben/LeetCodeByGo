package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type ListNode = common.ListNode

func main() {
	str := "1->2->3->4"
	fmt.Println("Output:", swapPairs(common.MakeListNode(str)))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.1 MB, 在所有 Go 提交中击败了22.48%的用户
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next != nil {
		node := head.Next
		node.Next, head.Next = head, swapPairs(node.Next)
		return node
	}
	return head
}
