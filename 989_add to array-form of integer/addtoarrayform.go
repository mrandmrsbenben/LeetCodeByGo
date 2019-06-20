package main

import (
	"fmt"
)

func main() {
	// A := []int{1, 2, 0, 0}
	// K := 34
	// A := []int{2, 7, 4}
	// K := 181
	// A := []int{2, 1, 5}
	// K := 806
	// A := []int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3}
	// K := 516
	A := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	K := 1
	// fmt.Println(math.Log10(float64(K)))
	fmt.Printf("Output: %v\n", addToArrayForm(A, K))
}

func addToArrayForm(A []int, K int) []int {
	arrK := make([]int, 0)
	for {
		if K < 10 {
			arrK = append(arrK, K)
			break
		} else {
			arrK = append(arrK, K%10)
			K = K / 10
		}
	}
	for i := 0; i < len(arrK)/2; i++ {
		arrK[i], arrK[len(arrK)-1-i] = arrK[len(arrK)-1-i], arrK[i]
	}
	if len(arrK) > len(A) {
		A, arrK = arrK, A
	}
	var add, sum int
	for i, j := len(A)-1, len(arrK)-1; i >= 0; i, j = i-1, j-1 {
		if j >= 0 {
			sum = A[i] + arrK[j] + add
			if sum > 9 {
				A[i] = sum % 10
				add = sum / 10
			} else {
				A[i] = sum
				add = 0
			}
		} else {
			sum = A[i] + add
			if sum > 9 {
				A[i] = sum % 10
				add = sum / 10
			} else {
				A[i] = sum
				add = 0
				break
			}
		}
	}
	if add > 0 {
		for {
			if add < 10 {
				A = append([]int{add}, A...)
				break
			} else {
				A = append([]int{add % 10}, A...)
				add = add / 10
			}
		}
	}
	return A
}

/*
func addToArrayForm(A []int, K int) []int {
	pow10 := func(a, b int) int64 {
		var p int64 = 1
		for i := 0; i < b; i++ {
			p *= 10
		}
		return p * int64(a)
	}
	var a int64
	for i := range A {
		a += pow10(A[i], len(A)-1-i)
	}
	sum := a + int64(K)
	fmt.Println(sum)
	suma := make([]int, 0)
	for {
		if sum < 10 {
			suma = append(suma, int(sum))
			break
		} else {
			suma = append(suma, int(sum%10))
			sum = sum / 10
		}
	}
	for i := 0; i < len(suma)/2; i++ {
		suma[i], suma[len(suma)-1-i] = suma[len(suma)-1-i], suma[i]
	}
	return suma
}*/
