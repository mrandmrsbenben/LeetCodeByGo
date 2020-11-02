package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// ListNode List Node
type ListNode = common.ListNode

func main() {
	// head := common.MakeListNode("4->2->1->3")
	head := common.MakeListNode("-1->5->3->4->0")
	fmt.Println("Output: ", insertionSortList(head))
}

//执行用时：4 ms, 在所有 Go 提交中击败了97.09%的用户
//内存消耗：3.2 MB, 在所有 Go 提交中击败了55.47%的用户
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var now, buf, buf2 *ListNode
	now = head
	for now.Next != nil {
		if now.Next.Val >= now.Val {
			now = now.Next
		} else {
			if head.Val > now.Next.Val {
				buf = now.Next
				now.Next = now.Next.Next
				buf.Next, head = head, buf
			} else {
				buf = head
				for buf.Next != nil {
					if buf.Next.Val > now.Next.Val {
						break
					}
					buf = buf.Next
				}
				buf2 = now.Next.Next
				buf.Next, now.Next.Next = now.Next, buf.Next
				now.Next = buf2
			}
		}
	}
	return head
}
