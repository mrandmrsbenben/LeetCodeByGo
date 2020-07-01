package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
	"strconv"
)

// TreeNode Tree Node
type TreeNode = common.TreeNode

func main() {
	S := "1-2--3--4-5--6--7"
	fmt.Println("Output: ", recoverFromPreorder(S))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :5 MB, 在所有 Go 提交中击败了100.00%的用户
func recoverFromPreorder(S string) *TreeNode {
	if len(S) == 0 {
		return nil
	}
	nodes := make([][]*TreeNode, 1)
	depth := 0
	start, val := -1, 0
	var cur, par *TreeNode
	for i := range S {
		if S[i] == '-' {
			if start != -1 {
				val, _ = strconv.Atoi(S[start:i])
				if len(nodes) <= depth {
					nodes = append(nodes, make([]*TreeNode, 1))
				}
				cur = &TreeNode{Val: val, Left: nil, Right: nil}
				nodes[depth] = append(nodes[depth], cur)

				//Find Parent Node
				if depth > 0 {
					par = nodes[depth-1][len(nodes[depth-1])-1]
					if par.Left == nil {
						par.Left = cur
					} else {
						par.Right = cur
					}
				}
				depth = 0
				start = -1
			}
			depth++
		} else if start == -1 {
			start = i
		}
	}
	if start != -1 {
		val, _ = strconv.Atoi(S[start:])
		if len(nodes) <= depth {
			nodes = append(nodes, make([]*TreeNode, 1))
		}
		cur = &TreeNode{Val: val, Left: nil, Right: nil}
		nodes[depth] = append(nodes[depth], cur)

		//Find Parent Node
		if depth > 0 {
			par = nodes[depth-1][len(nodes[depth-1])-1]
			if par.Left == nil {
				par.Left = cur
			} else {
				par.Right = cur
			}
		}
	}
	return nodes[0][0]
}
