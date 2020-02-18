package main

import "fmt"

func main() {
	fmt.Println("Output: ", sumZero(5))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.3 MB, 在所有 Go 提交中击败了71.05%的用户
func sumZero(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n/2; i++ {
		nums[i*2] = i + 1
		nums[i*2+1] = -1 * (i + 1)
	}
	if n%2 == 1 {
		nums[n-1] = 0
	}
	return nums
}
