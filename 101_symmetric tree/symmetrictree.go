package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
	"strconv"
)

type TreeNode = common.TreeNode

func main() {
	// vals := []int{1, 2, 2, 3, 4, 4, 3, 5, 7, -1, 6, 6, -1, 7, 5}
	// vals := []int{1, 2, 2, 2, -1, -1, 2}
	// vals := []int{1, 2, 2, -1, 3, -1, 3}
	vals := []int{5, 2, 2, -1, 2, -1, 2, -1, -1, 2, -1, -1, -1, 2, -1}
	fmt.Printf("Output: %v\n", isSymmetric(common.MakeTree(vals)))
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var lf, rf func(t *TreeNode)
	lleaves := ""
	lf = func(t *TreeNode) {
		if t == nil {
			return
		}
		if t.Left == nil && t.Right == nil {
			lleaves = lleaves + strconv.Itoa(t.Val) + "->"
			return
		}
		if t.Left != nil {
			lleaves = lleaves + "L->"
			lf(t.Left)
		} else {
			lleaves = lleaves + "null->"
		}
		lleaves = lleaves + strconv.Itoa(t.Val) + "->"
		if t.Right != nil {
			lleaves = lleaves + "R->"
			lf(t.Right)
		} else {
			lleaves = lleaves + "null->"
		}
	}
	rleaves := ""
	rf = func(t *TreeNode) {
		if t == nil {
			return
		}
		if t.Left == nil && t.Right == nil {
			rleaves = rleaves + strconv.Itoa(t.Val) + "->"
			return
		}
		if t.Right != nil {
			rleaves = rleaves + "L->"
			rf(t.Right)
		} else {
			rleaves = rleaves + "null->"
		}
		rleaves = rleaves + strconv.Itoa(t.Val) + "->"
		if t.Left != nil {
			rleaves = rleaves + "R->"
			rf(t.Left)
		} else {
			rleaves = rleaves + "null->"
		}
	}
	lf(root.Left)
	rf(root.Right)
	fmt.Println(lleaves)
	fmt.Println(rleaves)
	return lleaves == rleaves
}
