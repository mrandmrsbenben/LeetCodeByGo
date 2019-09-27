package main

import "fmt"

func main() {
	// cont := []int{3, 2, 0, 2}
	// cont := []int{0, 0, 3}
	// cont := []int{0, 2}
	cont := []int{2}
	fmt.Println("Output: ", fraction(cont))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func fraction(cont []int) []int {
	n, m := f(cont)
	return []int{n, m}
}

func f(cont []int) (int, int) {
	if len(cont) == 1 {
		return cont[0], 1
	}
	n := cont[0]
	x, y := f(cont[1:])
	n = n*x + y
	return n, x
}
