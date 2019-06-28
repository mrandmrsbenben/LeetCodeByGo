package main

import "fmt"

func main() {
	k, n := 3, 9
	fmt.Println("Output:", combinationSum3(k, n))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :9.9 MB, 在所有 Go 提交中击败了11.11%的用户
func combinationSum3(k int, n int) [][]int {
	if k >= n || n > 45 {
		return [][]int{}
	}
	combs := make([][]int, 0)
	var combSum func(k, n int, comb []int)
	combSum = func(k, n int, comb []int) {
		if k == 1 {
			if n > comb[len(comb)-1] && n <= 9 {
				buf := make([]int, len(comb))
				copy(buf, comb)
				buf = append(buf, n)
				combs = append(combs, buf)
			}
		} else {
			for i := comb[len(comb)-1] + 1; i < n; i++ {
				buf := make([]int, len(comb))
				copy(buf, comb)
				buf = append(buf, i)
				combSum(k-1, n-i, buf)
			}
		}
	}
	for i := 1; i < n; i++ {
		if i >= 9 {
			break
		}
		comb := []int{i}
		combSum(k-1, n-i, comb)
	}
	return combs
}
