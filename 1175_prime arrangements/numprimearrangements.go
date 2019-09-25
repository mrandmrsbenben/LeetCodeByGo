package main

import "fmt"

func main() {
	fmt.Println("Output: ", numPrimeArrangements(100))
}

func numPrimeArrangements(n int) int {
	mod := 1000000007
	ps := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	count := 0
	for i := 0; i < len(ps); i++ {
		if ps[i] <= n {
			count++
		} else {
			break
		}
	}

	return f(count) * f(n-count) % mod
}

func f(x int) int {
	sum := 1
	mod := 1000000007
	for i := 2; i <= x; i++ {
		sum = (sum * i) % mod
	}
	return sum
}
