package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Output: ", sortByBits(arr))
}

//执行用时 :24 ms, 在所有 Go 提交中击败了20.00%的用户
//内存消耗 :6.1 MB, 在所有 Go 提交中击败了100.00%的用户
func sortByBitsV1(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		ci := strings.Count(strconv.FormatInt(int64(arr[i]), 2), "1")
		cj := strings.Count(strconv.FormatInt(int64(arr[j]), 2), "1")
		if ci == cj {
			return arr[i] < arr[j]
		}
		return ci < cj
	})
	return arr
}

//执行用时 :8 ms, 在所有 Go 提交中击败了83.33%的用户
//内存消耗 :4.5 MB, 在所有 Go 提交中击败了100.00%的用户
func sortByBits(arr []int) []int {
	bits := make([][]int, 14)
	cnt := 0
	for _, n := range arr {
		i := strings.Count(strconv.FormatInt(int64(n), 2), "1")
		bits[i] = append(bits[i], n)
		cnt++
	}

	ret := make([]int, cnt)
	k := 0
	for i := range bits {
		sort.Ints(bits[i])
		for j := range bits[i] {
			ret[k] = bits[i][j]
			k++
		}
	}

	return ret
}
