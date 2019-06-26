package main

import (
	common "LeetCode/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	fmt.Println("Output:", allPossibleFBT(4))
}

//执行用时 :1196 ms, 在所有 Go 提交中击败了50.00%的用户
//内存消耗 :205 MB, 在所有 Go 提交中击败了33.33%的用户
func allPossibleFBT(N int) []*TreeNode {
	trees := make([]*TreeNode, 0)
	if N%2 == 0 {
		return trees
	} else if N == 1 {
		trees = append(trees, &TreeNode{0, nil, nil})
		return trees
	}
	N = N - 1
	i := 1
	for N > i {
		left := allPossibleFBT(i)
		right := allPossibleFBT(N - i)
		nodes := make([]*TreeNode, len(left)*len(right))
		index := 0
		for _, l := range left {
			for _, r := range right {
				nodes[index] = &TreeNode{0, l, r}
				index++
			}
		}
		i += 2
		trees = append(trees, nodes...)
	}
	return trees
}
