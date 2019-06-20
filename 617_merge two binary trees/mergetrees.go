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
	fmt.Println(makeTree1())
	fmt.Println(makeTree2())
	fmt.Println(mergeTrees(makeTree1(), makeTree2()))
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 != nil && t2 == nil {
		return t1
	} else if t1 == nil && t2 != nil {
		return t2
	} else if t1 == nil && t2 == nil {
		return nil
	}

	t := &TreeNode{t1.Val + t2.Val, nil, nil}

	if t1.Left == nil && t2.Left != nil {
		t.Left = t2.Left
	} else if t1.Left != nil && t2.Left == nil {
		t.Left = t1.Left
	} else if t1.Left != nil && t2.Left != nil {
		t.Left = mergeTrees(t1.Left, t2.Left)
	}

	if t1.Right == nil && t2.Right != nil {
		t.Right = t2.Right
	} else if t1.Right != nil && t2.Right == nil {
		t.Right = t1.Right
	} else if t1.Right != nil && t2.Right != nil {
		t.Right = mergeTrees(t1.Right, t2.Right)
	}
	return t
}

func makeTree1() *TreeNode {
	t := &TreeNode{1, nil, nil}
	t.Left = &TreeNode{3, nil, nil}
	t.Right = &TreeNode{2, nil, nil}
	t.Left.Left = &TreeNode{5, nil, nil}
	return t
}

func makeTree2() *TreeNode {
	t := &TreeNode{2, nil, nil}
	t.Left = &TreeNode{1, nil, nil}
	t.Right = &TreeNode{3, nil, nil}
	t.Left.Right = &TreeNode{4, nil, nil}
	t.Right.Right = &TreeNode{7, nil, nil}
	return t
}
