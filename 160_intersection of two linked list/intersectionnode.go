package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type ListNode = common.ListNode

func main() {
	strA := "4"
	strB := "4"
	fmt.Printf("Ouptput: %v\n",
		getIntersectionNode(common.MakeListNode(strA), common.MakeListNode(strB)))
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	arrA := listToArray(headA)
	arrB := listToArray(headB)
	if len(arrB) > len(arrA) {
		arrA, arrB = arrB, arrA
	}
	var prev, cur *ListNode
	for i, j := len(arrA)-1, len(arrB)-1; i >= 0 && j >= 0; i-- {
		if arrA[i] == arrB[j] {
			cur = &ListNode{Val: arrA[i], Next: prev}
			prev = cur
			j--
		} else if arrA[i] == 0 || arrB[j] == 0 {
			return prev.Next
		} else {
			return prev
		}
	}
	return cur
}

func listToArray(head *ListNode) []int {
	arr := make([]int, 0)
	for {
		arr = append(arr, head.Val)
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	return arr
}
