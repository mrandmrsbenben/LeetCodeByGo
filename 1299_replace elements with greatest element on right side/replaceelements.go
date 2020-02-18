package main

import "fmt"

func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Println("Output: ", replaceElements(arr))
}

//执行用时 :16 ms, 在所有 Go 提交中击败了93.42%的用户
//内存消耗 :6 MB, 在所有 Go 提交中击败了58.06%的用户
func replaceElements(arr []int) []int {
	i := len(arr) - 1
	max := 0
	for i > 0 {
		if arr[i] < max {
			arr[i] = max
		} else {
			max = arr[i]
		}
		i--
	}
	return append(arr[1:], -1)
}
