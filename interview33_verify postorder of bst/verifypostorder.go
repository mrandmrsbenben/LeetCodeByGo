package main

import "fmt"

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
func main() {
	postorder := []int{1, 6, 3, 2, 5}
	// postorder := []int{5, 4, 3, 2, 1}
	fmt.Println("Output: ", verifyPostorder(postorder))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.1 MB, 在所有 Go 提交中击败了100.00%的用户
func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	var i int
	root := postorder[len(postorder)-1]
	postorder = postorder[0 : len(postorder)-1]
	for i = len(postorder) - 1; i >= 0; i-- {
		if postorder[i] < root {
			break
		}
	}
	var l, r []int
	if i < 0 {
		l = []int{}
		r = postorder
	} else {
		l = postorder[0 : i+1]
		r = postorder[i+1:]
	}
	for _, n := range l {
		if n > root {
			return false
		}
	}
	return verifyPostorder(l) && verifyPostorder(r)
}
