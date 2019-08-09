package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	root := common.MakeTree([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Println("Output:", zigzagLevelOrder(root))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :3.3 MB, 在所有 Go 提交中击败了20.90%的用户
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}
	allVals := make([][]int, 0)
	nextLevel := []*TreeNode{root}
	var nodes []*TreeNode
	var vals []int
	levelCount := 1
	for {
		nodes = make([]*TreeNode, 0)
		vals = make([]int, len(nextLevel))
		if levelCount%2 == 1 {
			for i, node := range nextLevel {
				vals[i] = node.Val
				if node.Left != nil {
					nodes = append(nodes, node.Left)
				}
				if node.Right != nil {
					nodes = append(nodes, node.Right)
				}
			}
		} else {
			for i, node := range nextLevel {
				vals[len(nextLevel)-1-i] = node.Val
				if node.Left != nil {
					nodes = append(nodes, node.Left)
				}
				if node.Right != nil {
					nodes = append(nodes, node.Right)
				}
			}
		}
		allVals = append(allVals, vals)
		if len(nodes) == 0 {
			break
		}
		nextLevel = make([]*TreeNode, len(nodes))
		copy(nextLevel, nodes)
		levelCount++
	}
	return allVals
}
