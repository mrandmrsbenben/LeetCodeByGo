package main

import "fmt"

func main() {
	A := []int{1, 2, 3, 4}
	queries := [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}
	fmt.Printf("Output: %v\n", sumEvenAfterQueries(A, queries))
}

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	sum := make([]int, len(A))
	orgs := 0
	for _, n := range A {
		if n%2 == 0 {
			orgs = orgs + n
		}
	}
	for i := range queries {
		if A[queries[i][1]]%2 == 0 {
			orgs = orgs - A[queries[i][1]]
		}
		A[queries[i][1]] = A[queries[i][1]] + queries[i][0]
		if A[queries[i][1]]%2 == 0 {
			orgs = orgs + A[queries[i][1]]
		}
		sum[i] = orgs
	}
	return sum
}
