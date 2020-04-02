package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

//ListNode Linked List Node
type ListNode = common.ListNode

// https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
func main() {
	headA := common.MakeListNodeByArray([]int{4, 1, 8, 4, 5})
	headB := common.MakeListNodeByArray([]int{5, 0, 1, 8, 4, 5})
	fmt.Println("Output: ", getIntersectionNode(headA, headB))
}

//执行用时 :40 ms, 在所有 Go 提交中击败了98.77%的用户
//内存消耗 :8.4 MB, 在所有 Go 提交中击败了100.00%的用户
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}
		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	return a
}
