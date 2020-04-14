package main

import "fmt"

func main() {
	// nums := []int{0, 0, 1, 0, 0, 0, 1, 1}
	nums := []int{}
	fmt.Println("Output: ", findMaxLength(nums))
}

//执行用时 :120 ms, 在所有 Go 提交中击败了95.12%的用户
//内存消耗 :7.1 MB, 在所有 Go 提交中击败了100.00%的用户
func findMaxLength(nums []int) int {
	max, cnt := 0, 0
	cm := make(map[int]int)
	for i := range nums {
		if nums[i] == 0 {
			cnt--
		} else {
			cnt++
		}
		if cnt == 0 && max < i+1 {
			max = i + 1
		}
		if index, ok := cm[cnt]; ok {
			if max < i-index {
				max = i - index
			}
		} else {
			cm[cnt] = i
		}
	}
	return max
}
