package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 6, 6, 6, 6, 7, 10}
	fmt.Println("Output: ", findSpecialInteger(arr))
}

//执行用时 :12 ms, 在所有 Go 提交中击败了82.81%的用户
//内存消耗 :5.3 MB, 在所有 Go 提交中击败了78.95%的用户
func findSpecialInteger(arr []int) int {
	n := len(arr) / 4
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] == arr[i+n] {
			return arr[i]
		}
		if arr[len(arr)-1-i] == arr[len(arr)-1-i-n] {
			return arr[len(arr)-1-i]
		}
	}
	return arr[0]
}
