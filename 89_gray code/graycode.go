package main

import (
	"fmt"
)

func main() {
	fmt.Println("Output:", grayCode(3))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了94.59%的用户
//内存消耗 :7.7 MB, 在所有 Go 提交中击败了14.64%的用户
func grayCode(n int) []int {
	ns := []int{0, 1}
	switch n {
	case 0:
		return []int{0}
	case 1:
		return ns
	}
	bit := 2
	for n > 1 {
		as := make([]int, len(ns))
		j := len(as) - 1
		for i := range ns {
			as[j] = ns[i] + bit
			j--
		}
		ns = append(ns, as...)
		n--
		bit *= 2
	}
	return ns
}
