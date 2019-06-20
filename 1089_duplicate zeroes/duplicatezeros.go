package main

import "fmt"

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr)
	fmt.Println("Output:", arr)
}

// Runtime: 36 ms, faster than 52.78% of Go online submissions for Duplicate Zeros.
// Memory Usage: 7.8 MB, less than 100.00% of Go online submissions for Duplicate Zeros.
func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 && i < len(arr)-1 {
			for j := len(arr) - 1; j > i+1; j-- {
				arr[j] = arr[j-1]
			}
			arr[i+1] = 0
			i++
		}
	}
}

// Runtime: 44 ms, faster than 25.00% of Go online submissions for Duplicate Zeros.
// Memory Usage: 8.1 MB, less than 100.00% of Go online submissions for Duplicate Zeros.
func duplicateZerosV1(arr []int) {
	var buf []int
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 && i < len(arr)-1 {
			buf = make([]int, len(arr)-i-1)
			copy(buf, arr[i+1:i+len(buf)])
			arr[i+1] = 0
			for j := 0; j < len(buf)-1; j++ {
				arr[i+2+j] = buf[j]
			}
			i++
		}
	}
}
