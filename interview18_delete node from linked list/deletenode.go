package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// ListNode linked-list node
type ListNode = common.ListNode

func main() {
	head := common.MakeListNodeByArray([]int{4, 5, 1, 9})
	fmt.Println("Output: ", deleteNode(head, 9))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.9 MB, 在所有 Go 提交中击败了100.00%的用户
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	} else if head.Val == val {
		return head.Next
	}

	prev, cur := head, head.Next
	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
			break
		}
		prev, cur = cur, cur.Next
	}
	return head
}
