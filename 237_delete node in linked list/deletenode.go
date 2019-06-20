package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	str := ""
	for {
		if len(str) != 0 {
			str = str + "->"
		}
		str = str + strconv.Itoa(node.Val)
		node = node.Next
		if node == nil {
			break
		}
	}
	return str
}

func main() {
	node0 := &ListNode{9, nil}
	node1 := &ListNode{1, node0}
	node2 := &ListNode{5, node1}
	node3 := &ListNode{4, node2}
	fmt.Println("Before Delete:", node3)
	deleteNode(node3)
	fmt.Println("After Delete:", node3)
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
