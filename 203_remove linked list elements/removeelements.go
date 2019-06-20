package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type ListNode = common.ListNode

func main() {
	str := "1->1"
	val := 1
	fmt.Printf("Output: %v\n", removeElements(common.MakeListNode(str), val))
}

// 执行用时 : 8 ms, 在Remove Linked List Elements的Go提交中击败了98.80% 的用户
// 内存消耗 : 5.4 MB, 在Remove Linked List Elements的Go提交中击败了6.12% 的用户
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		if head.Val == val {
			return head.Next
		}
		return head
	} else {
		if val == head.Next.Val {
			head.Next = head.Next.Next
			removeElements(head, val)
		} else {
			removeElements(head.Next, val)
		}
	}
	if head.Val == val {
		return head.Next
	}
	return head
}
