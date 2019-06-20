package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{3, 9, 20, -1, -1, 15, 7}
	root := common.MakeTree(vals)
	fmt.Println(root)
	fmt.Printf("Output: %v\n", levelOrderBottom(root))
}

func levelOrderBottom(root *TreeNode) [][]int {
	rtn := make([][]int, 0)
	if root == nil {
		return rtn
	}

	vals := [][]int{{root.Val}}
	nodes := [][]*TreeNode{{root}}
	cnt, level := 1, 1
	levelNodes := make([]*TreeNode, 0)
	levelVals := make([]int, 0)
	for {
		if root.Left != nil {
			levelNodes = append(levelNodes, root.Left)
			levelVals = append(levelVals, root.Left.Val)
		}
		if root.Right != nil {
			levelNodes = append(levelNodes, root.Right)
			levelVals = append(levelVals, root.Right.Val)
		}
		if cnt >= len(nodes[level-1])-1 {
			cnt = 0
			level = level + 1
			if len(levelNodes) == 0 {
				break
			}
			nodes = append(nodes, levelNodes)
			vals = append(vals, levelVals)
			levelNodes = make([]*TreeNode, 0)
			levelVals = make([]int, 0)
		} else {
			cnt = cnt + 1
		}
		root = nodes[level-1][cnt]
	}
	for i := 0; i < len(vals)/2; i++ {
		vals[i], vals[len(vals)-1-i] = vals[len(vals)-1-i], vals[i]
	}
	return vals
}
