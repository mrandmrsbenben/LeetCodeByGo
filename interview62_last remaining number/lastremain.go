package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
func main() {
	// fmt.Println("Output: ", lastRemaining(5, 3))
	// fmt.Println("Output: ", lastRemaining(10, 17))
	fmt.Println("Output: ", lastRemaining(70866, 116922))
}

//执行用时 :8 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func lastRemaining(n int, m int) int {
	res := 0
	for i := 2; i <= n; i++ {
		res = (m + res) % i
	}
	return res
}

//执行用时 :24 ms, 在所有 Go 提交中击败了14.00%的用户
//内存消耗 :7.8 MB, 在所有 Go 提交中击败了100.00%的用户
func lastRemainingV1(n int, m int) int {
	if n == 0 {
		return 0
	}
	x := lastRemaining(n-1, m)
	return (m + x) % n
}

func lastRemainingSlowV2(n int, m int) int {
	numbers := make([]int, n)
	for i := range numbers {
		numbers[i] = i
	}
	index := 0
	for n > 1 {
		if n < m {
			index = (m - 1) % n
		} else {
			index = m - 1
		}
		left := numbers[0:index]
		right := numbers[index+1:]
		numbers = append(right, left...)
		n = len(numbers)
	}

	return numbers[0]
}

func lastRemainingSlow(n int, m int) int {
	var head, prev, tail *ListNode

	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			tail = &ListNode{i, nil}
			prev = tail
		} else {
			head = &ListNode{i, prev}
			prev = head
		}
	}
	tail.Next = head
	cnt := 1
	for n > 1 {
		if cnt%m == 0 {
			fmt.Println(head.Val)
			head.Val = head.Next.Val
			head.Next = head.Next.Next
			n--
		} else {
			head = head.Next
		}
		cnt++
	}
	return head.Val
}

// ListNode Linked List Node
type ListNode struct {
	Val  int
	Next *ListNode
}
