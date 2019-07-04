package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Println("Output:", singleNumber(nums))
}

//执行用时 :8 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :5.1 MB, 在所有 Go 提交中击败了17.65%的用户
func singleNumber(nums []int) []int {
	m := make(map[int]int, len(nums))
	for _, n := range nums {
		if _, ok := m[n]; ok {
			m[n] = 1
		} else {
			m[n] = 0
		}
	}
	ns := make([]int, 2)
	i := 0
	for n := range m {
		if m[n] == 0 {
			ns[i] = n
			i++
		}
	}
	return ns
}
