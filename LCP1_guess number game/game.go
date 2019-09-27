package main

import "fmt"

func main() {
	guess := []int{1, 2, 3}
	answer := []int{1, 2, 3}
	fmt.Println("Output: ", game(guess, answer))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func game(guess []int, answer []int) int {
	cnt := 0
	for i := range guess {
		if guess[i] == answer[i] {
			cnt++
		}
	}
	return cnt
}
