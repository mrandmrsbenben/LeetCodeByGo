package main

import "fmt"

type testcase struct {
	input  [][]int
	output [][]int
}

func main() {
	// t := &testcase{[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
	// 	[][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}}
	// t := &testcase{[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
	// 	[][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}}
	t := &testcase{[][]int{{0}, {0}}, [][]int{{1}, {1}}}
	output := flipAndInvertImage(t.input)
	for i := 0; i < len(output); i++ {
		for j := 0; j < len(output[i]); j++ {
			if output[i][j] != t.output[i][j] {
				fmt.Printf("Fail.Expect:%v, Actual:%v", t.output, output)
				break
			}
		}
	}
	fmt.Println(output)
}

func flipAndInvertImage(A [][]int) [][]int {
	for i := range A {
		for j := 0; j < len(A[i]); j++ {
			if j < len(A[i])/2 {
				A[i][j], A[i][len(A[i])-j-1] = A[i][len(A[i])-j-1], A[i][j]
			}
			A[i][j] = (A[i][j] + 1) % 2
		}
	}
	return A
}
