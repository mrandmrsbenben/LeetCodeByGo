package main

import "fmt"

func main() {
	// nums := []int{9, 1, 7, 9, 7, 9, 7}
	nums := []int{3, 4, 3, 3}
	fmt.Println("Output: ", singleNumber(nums))
}

//执行用时 :36 ms, 在所有 Go 提交中击败了43.56%的用户
//内存消耗 :6.4 MB, 在所有 Go 提交中击败了100.00%的用户
func singleNumber(nums []int) int {
	nm := make(map[int]int)
	res := 0
	for _, n := range nums {
		res += n
		nm[n]++
		if nm[n] == 3 {
			res -= n * 3
		}
	}
	return res
}

//执行用时 :64 ms, 在所有 Go 提交中击败了7.36%的用户
//内存消耗 :6.2 MB, 在所有 Go 提交中击败了100.00%的用户
func singleNumberV1(nums []int) int {
	bits := make([]int, 32)
	for _, n := range nums {
		mask := 1
		for j := 31; j >= 0; j-- {
			if n&mask != 0 {
				bits[j]++
			}
			mask = mask << 1
		}
	}
	res := 0
	for i := range bits {
		res = res << 1
		res += bits[i] % 3
	}
	return res
}
