package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type ListNode = common.ListNode

func main() {
	str := "1->2->2->1"
	fmt.Printf("Output: %v\n", isPalindrome(common.MakeListNode(str)))
}

// 执行用时 :24 ms, 在所有Go提交中击败了95.07%的用户
// 内存消耗 :6.5 MB, 在所有Go提交中击败了34.43%的用户
func isPalindrome(head *ListNode) bool {
	vals := make([]int, 0)
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	for i := 0; i < len(vals)/2; i++ {
		if vals[i] != vals[len(vals)-1-i] {
			return false
		}
	}
	return true
}
