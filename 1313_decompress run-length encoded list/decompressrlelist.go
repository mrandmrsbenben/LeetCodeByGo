package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println("Output: ", decompressRLElist(nums))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了87.04%的用户
//内存消耗 :4.4 MB, 在所有 Go 提交中击败了100.00%的用户
func decompressRLElist(nums []int) []int {
	cnt := 0
	for i := 0; i < len(nums); i += 2 {
		cnt += nums[i]
	}
	res := make([]int, cnt)
	k := 0
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			res[k] = nums[i+1]
			k++
		}
	}
	return res
}

//执行用时 :8 ms, 在所有 Go 提交中击败了39.81%的用户
//内存消耗 :5.7 MB, 在所有 Go 提交中击败了100.00%的用户
func decompressRLElistV2(nums []int) []int {
	res := []int{}
	var arr []int
	for i := 0; i < len(nums); i += 2 {
		arr = make([]int, nums[i])
		for j := 0; j < nums[i]; j++ {
			arr[j] = nums[i+1]
		}
		res = append(res, arr...)
	}

	return res
}

//执行用时 :8 ms, 在所有 Go 提交中击败了39.81%的用户
//内存消耗 :6 MB, 在所有 Go 提交中击败了100.00%的用户
func decompressRLElistV1(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			res = append(res, nums[i+1])
		}
	}
	return res
}
