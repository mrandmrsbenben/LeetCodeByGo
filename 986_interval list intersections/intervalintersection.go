package main

import "fmt"

func main() {
	A := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	B := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	// A := [][]int{{3, 5}, {9, 20}}
	// B := [][]int{{4, 5}, {7, 10}, {11, 12}, {14, 15}, {16, 20}}
	// A := [][]int{{0, 5}, {12, 14}, {15, 18}}
	// B := [][]int{{11, 15}, {18, 19}}
	// A := [][]int{{2, 3}, {4, 6}, {7, 9}, {10, 16}, {17, 20}}
	// B := [][]int{{1, 3}, {4, 7}, {10, 11}, {15, 18}}
	fmt.Println("Output: ", intervalIntersection(A, B))
}

//执行用时 :20 ms, 在所有 Go 提交中击败了92.45%的用户
//内存消耗 :6.3 MB, 在所有 Go 提交中击败了100.00%的用户
func intervalIntersection(A [][]int, B [][]int) [][]int {
	C := make([][]int, 0)

	i, j := 0, 0
	for i < len(A) && j < len(B) {
		if A[i][0] < B[j][0] {
			if A[i][1] >= B[j][0] {
				if A[i][1] > B[j][1] {
					C = append(C, []int{B[j][0], B[j][1]})
					A[i][0] = B[j][1] + 1
					j++
				} else {
					C = append(C, []int{B[j][0], A[i][1]})
					if A[i][1] == B[j][1] {
						j++
					} else {
						B[j][0] = A[i][1] + 1
					}
				}
			} else {
				i++
			}
		} else if A[i][0] > B[j][0] {
			if A[i][0] <= B[j][1] {
				if A[i][1] < B[j][1] {
					C = append(C, []int{A[i][0], A[i][1]})
					B[j][0] = A[i][1] + 1
					i++
				} else {
					C = append(C, []int{A[i][0], B[j][1]})
					if A[i][1] == B[j][1] {
						i++
					} else {
						A[i][0] = B[j][1] + 1
					}
				}
			} else {
				j++
			}
		} else {
			if A[i][1] < B[j][1] {
				C = append(C, []int{A[i][0], A[i][1]})
				B[j][0] = A[i][1] + 1
				i++
			} else if A[i][1] > B[j][1] {
				C = append(C, []int{A[i][0], B[j][1]})
				A[i][0] = B[j][1] + 1
				j++
			} else {
				C = append(C, []int{A[i][0], A[i][1]})
				i++
				j++
			}
		}
	}

	return C
}
