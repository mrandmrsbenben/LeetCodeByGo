package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// ListNode : Linked List Node
type ListNode = common.ListNode

// https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
func main() {
	vals := []int{1, 2, 3, 4, 5}
	head := common.MakeListNodeByArray(vals)
	fmt.Println("Output: ", getKthFromEnd(head, 2))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.2 MB, 在所有 Go 提交中击败了100.00%的用户
func getKthFromEnd(head *ListNode, k int) *ListNode {
	res := head
	for i := 1; head.Next != nil; i++ {
		if i >= k {
			res = res.Next
		}
		head = head.Next
	}
	return res
}
