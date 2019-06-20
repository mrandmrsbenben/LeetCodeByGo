package main

import "fmt"

func main() {
	s := []int{3, 1, 2, 4}
	fmt.Printf("Input:%d, Output:%d\n", s, sortArrayByParity(s))
}

func sortArrayByParity(A []int) []int {
	even := make([]int, 0)
	odd := make([]int, 0)
	for _, v := range A {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	return append(even, odd...)
}
