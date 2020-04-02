package main

import "fmt"

// https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
func main() {
	nums := []int{7, 5, 6, 4}
	fmt.Println("Output: ", reversePairs(nums))
}

//执行用时 :120 ms, 在所有 Go 提交中击败了85.93%的用户
//内存消耗 :8.9 MB, 在所有 Go 提交中击败了100.00%的用户
func reversePairs(nums []int) int {
	res := 0

	var merge func(l, r []int) []int
	merge = func(l, r []int) []int {
		ll, lr := len(l), len(r)
		mer := make([]int, ll+lr)

		var i, j, k int
		for i < ll && j < lr {
			if l[i] <= r[j] {
				mer[k] = l[i]
				i++
			} else {
				mer[k] = r[j]
				j++
				res += ll - i
			}
			k++
		}
		for i < ll {
			mer[k] = l[i]
			i++
			k++
		}
		for j < lr {
			mer[k] = r[j]
			j++
			k++
		}
		return mer
	}

	var divide func(n []int) []int
	divide = func(n []int) []int {
		l := len(n)
		if l < 2 {
			return n
		}
		return merge(divide(n[0:l/2]), divide(n[l/2:]))
	}
	divide(nums)
	return res
}
