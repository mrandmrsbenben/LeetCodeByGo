package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	var str = ""

	if t.Left != nil {
		str = str + t.Left.String()
		if !strings.HasSuffix(str, ",") {
			str = str + ","
		}
	}
	str = str + strconv.Itoa(t.Val) + ","
	if t.Right != nil {
		str = str + t.Right.String()
		if !strings.HasSuffix(str, ",") {
			str = str + ","
		}
	}
	return str
}

func main() {
	t := makeTree()
	fmt.Printf("Input:%v, Output:%v,", t, invertTree(t))
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{root.Val, invertTree(root.Right), invertTree(root.Left)}
}

func makeTree() *TreeNode {
	t := &TreeNode{4, nil, nil}
	t.Left = &TreeNode{2, nil, nil}
	t.Right = &TreeNode{7, nil, nil}
	t.Left.Left = &TreeNode{1, nil, nil}
	t.Left.Right = &TreeNode{3, nil, nil}
	t.Right.Left = &TreeNode{6, nil, nil}
	t.Right.Right = &TreeNode{9, nil, nil}
	return t
}
