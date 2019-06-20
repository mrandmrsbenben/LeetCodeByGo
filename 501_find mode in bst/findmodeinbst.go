package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{1, -1, 2, -1, -1, 3}
	fmt.Printf("Output: %v\n", findMode(common.MakeTree(vals)))
}

func findMode(root *TreeNode) []int {
	modes := make([]int, 0)
	if root == nil {
		return modes
	}
	modeMap := make(map[int]int)
	var fModeCount func(t *TreeNode)
	fModeCount = func(t *TreeNode) {
		if t == nil {
			return
		}
		fModeCount(t.Left)
		if cnt, ok := modeMap[t.Val]; ok {
			modeMap[t.Val] = cnt + 1
		} else {
			modeMap[t.Val] = 1
		}
		fModeCount(t.Right)
	}
	fModeCount(root)
	max := 0
	for k, v := range modeMap {
		if v > max {
			modes = []int{k}
			max = v
		} else if v == max {
			modes = append(modes, k)
		}
	}
	return modes
}
