package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type ListNode = common.ListNode

func main() {
	// str := "1->2->3->4->NULL"
	str := "1->2->3->4->5->NULL"
	// str := "2->1->3->5->6->4->7->NULL"
	// str := "2->1->3"
	fmt.Println("Output:", oddEvenList(common.MakeListNode(str)))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了97.06%的用户
//内存消耗 :3.3 MB, 在所有 Go 提交中击败了83.87%的用户
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var ho, he *ListNode
	var odd, even *ListNode
	i := 1
	for head != nil {
		if i%2 == 0 {
			if even == nil {
				even = head
				he = even
			} else {
				even.Next = head
				even = even.Next
			}
			if head.Next == nil {
				odd.Next = he
				break
			}
		} else {
			if odd == nil {
				odd = head
				ho = odd
			} else {
				odd.Next = head
				odd = odd.Next
			}
			if head.Next == nil {
				even.Next = nil
				head.Next = he
				break
			}
		}
		head = head.Next
		i++
	}
	return ho
}
