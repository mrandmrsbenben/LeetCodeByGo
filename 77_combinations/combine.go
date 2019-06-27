package main

import "fmt"

func main() {
	n, k := 5, 5
	fmt.Println("Output:", combine(n, k))
}

//执行用时 :236 ms, 在所有 Go 提交中击败了84.44%的用户
//内存消耗 :48.4 MB, 在所有 Go 提交中击败了63.46%的用户
func combine(n int, k int) [][]int {
	if k > n {
		return [][]int{}
	} else if k == n {
		c := make([]int, n)
		for i := 0; i < n; i++ {
			c[i] = i + 1
		}
		return [][]int{c}
	}
	count := combineCount(n, k)
	cs := make([][]int, count)
	index := 0
	if k == 1 {
		for i := 1; i <= n; i++ {
			cs[index] = []int{i}
			index++
		}
	} else {
		for i := n; i > 1; i-- {
			ret := combine(i-1, k-1)
			for _, rc := range ret {
				cs[index] = make([]int, k)
				cs[index][0] = i
				for j := range rc {
					cs[index][j+1] = rc[j]
				}
				index++
			}
		}
	}
	return cs
}

func combineCount(n, k int) int {
	switch k {
	case 1:
		return n
	case 2:
		return n * (n - 1) / 2
	case 3:
		return n * (n - 1) * (n - 2) / 6
	case n:
		return 1
	default:
		return combineCount(n-1, k-1) + combineCount(n-1, k)
	}
}
