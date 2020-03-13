package main

import "fmt"

// https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println("Output: ", exchange(nums))
}

//执行用时 :16 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :6.4 MB, 在所有 Go 提交中击败了100.00%的用户
func exchange(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	lcnt, rcnt := 0, 0
	for _, n := range nums {
		if n%2 == 1 {
			res[lcnt] = n
			lcnt++
		} else {
			res[l-1-rcnt] = n
			rcnt++
		}
	}
	return res
}

//执行用时 :24 ms, 在所有 Go 提交中击败了91.87%的用户
//内存消耗 :6.6 MB, 在所有 Go 提交中击败了100.00%的用户
func exchangeV1(nums []int) []int {
	l := len(nums)
	if l == 1 {
		return nums
	}
	odd := make([]int, l)
	even := make([]int, l)
	ocnt := 0
	ecnt := 0
	for _, n := range nums {
		if n%2 == 1 {
			odd[ocnt] = n
			ocnt++
		} else {
			even[ecnt] = n
			ecnt++
		}
	}
	return append(odd[0:ocnt], even[0:ecnt]...)
}
