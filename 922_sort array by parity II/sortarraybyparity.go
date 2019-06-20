package main

import "fmt"

func main() {
	// A := []int{4, 2, 5, 7}
	// A := []int{648, 831, 560, 986, 192, 424, 997, 829, 897, 843}
	A := []int{2, 3, 1, 1, 4, 0, 0, 4, 3, 3}
	fmt.Printf("Output:%v\n", sortArrayByParityII(A))
}

func sortArrayByParityII(A []int) []int {
	for i := range A {
		if i%2 != A[i]%2 {
			for j := i + 1; j < len(A); j++ {
				if j%2 != A[j]%2 && i%2 != j%2 {
					A[i], A[j] = A[j], A[i]
					break
				}
			}
		}
	}
	return A
}

func sortArrayByParity(A []int) []int {
	odd := make([]int, 0)
	even := make([]int, 0)
	for _, v := range A {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	sorted := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		if i%2 == 0 {
			sorted[i] = even[i/2]
		} else {
			sorted[i] = odd[i/2]
		}
	}
	return sorted
}
