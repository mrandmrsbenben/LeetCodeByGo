// https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println("Output: ", constructArr(a))
}

//执行用时 :24 ms, 在所有 Go 提交中击败了97.99%的用户
//内存消耗 :8.2 MB, 在所有 Go 提交中击败了100.00%的用户
func constructArr(a []int) []int {
	l, r := make([]int, len(a)), make([]int, len(a))

	m := 1
	l[0] = 1
	for i := 1; i < len(l); i++ {
		m *= a[i-1]
		l[i] = m
	}

	m = 1
	r[len(r)-1] = 1
	for i := len(r) - 2; i >= 0; i-- {
		m *= a[i+1]
		r[i] = m
	}

	res := make([]int, len(a))
	for i := range res {
		res[i] = l[i] * r[i]
	}
	return res
}
