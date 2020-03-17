package main

import "fmt"

// https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/
func main() {
	pushed := []int{1, 2, 3, 4, 5, 6, 7}
	popped := []int{4, 5, 3, 2, 1, 6, 7}
	// pushed := []int{1, 2, 3, 4, 5}
	// popped := []int{4, 3, 5, 1, 2}
	// pushed := []int{1, 2}
	// popped := []int{2, 1}
	fmt.Println("Output: ", validateStackSequences(pushed, popped))
}

//执行用时 :8 ms, 在所有 Go 提交中击败了88.33%的用户
//内存消耗 :3.6 MB, 在所有 Go 提交中击败了100.00%的用户
func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) <= 1 {
		return true
	}
	var wait []int
	for i := range pushed {
		if popped[0] == pushed[i] {
			wait = pushed[i+1:]
			pushed = pushed[0:i]
			break
		}
	}
	popped = popped[1:]
	for i := range popped {
		if len(pushed) > 0 && popped[i] == pushed[len(pushed)-1] {
			pushed = pushed[0 : len(pushed)-1]
		} else {
			for j := range wait {
				if popped[i] == wait[j] {
					pushed = append(pushed, wait[0:j]...)
					wait = wait[j+1:]
					break
				}
			}
		}
	}
	return len(pushed)+len(wait) == 0
}
