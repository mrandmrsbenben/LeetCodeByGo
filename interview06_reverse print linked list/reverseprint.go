package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// ListNode Linked List Node
type ListNode = common.ListNode

func main() {
	head := common.MakeListNodeByArray([]int{1, 3, 2})
	fmt.Println("Output: ", reversePrint(head))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :3.1 MB, 在所有 Go 提交中击败了100.00%的用户
func reversePrint(head *ListNode) []int {
	ret := make([]int, 0)
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	l := len(ret)
	for i := 0; i < l/2; i++ {
		ret[i], ret[l-1-i] = ret[l-1-i], ret[i]
	}
	return ret
}
