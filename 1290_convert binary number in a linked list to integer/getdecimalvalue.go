package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// ListNode From Common
type ListNode = common.ListNode

func main() {
	l := common.MakeListNodeByArray([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0})
	fmt.Println("Output: ", getDecimalValue(l))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2 MB, 在所有 Go 提交中击败了60.45%的用户
func getDecimalValue(head *ListNode) int {
	sum := 0
	for head != nil {
		sum = sum*2 + head.Val
		head = head.Next
	}
	return sum
}
