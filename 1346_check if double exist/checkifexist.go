package main

import "fmt"

func main() {
	arr := []int{10, 2, 5, 3}
	fmt.Println("Output: ", checkIfExist(arr))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :3.3 MB, 在所有 Go 提交中击败了100.00%的用户
func checkIfExist(arr []int) bool {
	marr := make(map[int]int)
	for _, n := range arr {
		if marr[n*2] == 1 {
			return true
		} else if n%2 == 0 && marr[n/2] == 1 {
			return true
		}
		marr[n] = 1
	}
	return false
}
