package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	A := []int{1, 1}
	B := []int{2, 2}
	fmt.Printf("Output: %v\n", fairCandySwap(A, B))
}

func fairCandySwap(A []int, B []int) []int {
	swap := make([]int, 2)
	sort.Ints(A)
	sort.Ints(B)
	var suma, sumb int
	for _, v := range A {
		suma = suma + v
	}
	for _, v := range B {
		sumb = sumb + v
	}
	if suma > sumb {
		A, B = B, A
	}
	var i int
	n := int(math.Abs(float64(suma-sumb))) / 2
	for _, v := range A {
		i = sort.SearchInts(B, v+n)
		if i < len(B) && B[i] == v+n {
			swap[0] = v
			swap[1] = B[i]
			break
		}
	}
	if suma > sumb {
		swap[0], swap[1] = swap[1], swap[0]
	}
	return swap
}
