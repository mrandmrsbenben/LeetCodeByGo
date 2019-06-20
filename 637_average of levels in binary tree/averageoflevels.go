package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{3, 9, 20, -1, -1, 15, 7}
	root := common.MakeTree(vals)
	fmt.Printf("Output: %v\n", averageOfLevels(root))
}

func averageOfLevels(root *TreeNode) []float64 {
	vals := make([]float64, 1)
	if root == nil {
		return vals
	}
	vals[0] = float64(root.Val)
	nodes := [][]*TreeNode{{root}}
	cnt, level := 1, 1
	levelSum := 0.0
	levelNodes := make([]*TreeNode, 0)
	for {
		if root.Left != nil {
			levelNodes = append(levelNodes, root.Left)
			levelSum = levelSum + float64(root.Left.Val)
		}
		if root.Right != nil {
			levelNodes = append(levelNodes, root.Right)
			levelSum = levelSum + float64(root.Right.Val)
		}
		if cnt >= len(nodes[level-1])-1 {
			cnt = 0
			level = level + 1
			if len(levelNodes) == 0 {
				break
			}
			nodes = append(nodes, levelNodes)
			vals = append(vals, levelSum/float64(len(levelNodes)))
			levelNodes = make([]*TreeNode, 0)
			levelSum = 0
		} else {
			cnt = cnt + 1
		}
		root = nodes[level-1][cnt]
	}
	return vals
}
