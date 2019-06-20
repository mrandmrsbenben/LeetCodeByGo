package common

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	if node == nil {
		return "NULL"
	}
	str := strconv.Itoa(node.Val) + "->"
	if node.Next == nil {
		str = str + "NULL"
	} else {
		str = str + node.Next.String()
	}
	return str
}

func MakeListNode(str string) *ListNode {
	index := strings.Index(str, "->")
	if index != -1 {
		s := str[0:index]
		val, err := strconv.Atoi(s)
		if err == nil {
			return &ListNode{val, MakeListNode(str[index+2:])}
		}
	} else {
		val, err := strconv.Atoi(str)
		if err == nil {
			return &ListNode{val, nil}
		}
	}
	return nil
}

func MakeListNodeByArray(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := []*ListNode{}
	for _, v := range vals {
		node := &ListNode{v, nil}
		if len(nodes) > 0 {
			nodes[len(nodes)-1].Next = node
		}
		nodes = append(nodes, node)
	}
	return nodes[0]
}
