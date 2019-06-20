package main

import "fmt"

func main() {
	A := []string{"cba", "daf", "ghi"}
	fmt.Println("Output:", minDeletionSize(A))
}

func minDeletionSize(A []string) int {
	size := 0
	if len(A) == 0 {
		return 0
	}
	l := len(A[0])
	var s1, s2 string
	for i := 0; i < l; i++ {
		for j := 1; j < len(A); j++ {
			s1 = A[j-1][i : i+1]
			s2 = A[j][i : i+1]
			if s2 < s1 {
				size = size + 1
				break
			}
		}
	}
	return size
}
