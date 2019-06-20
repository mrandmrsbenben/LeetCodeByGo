package main

import "fmt"

type testcases struct {
	input  []int
	expect int
}

func main() {
	cases := []testcases{{[]int{1, 2, 3, 4}, 3},
		{[]int{2, 1, 2, 5, 3, 2}, 2},
		{[]int{5, 1, 5, 2, 5, 3, 5, 4}, 5}}
	for _, t := range cases {
		fmt.Printf("Input: %v, Expect:%d, Actual:%d\n",
			t.input, t.expect, repeatedNTimes(t.input))
	}
}

func repeatedNTimes(A []int) int {
	n := len(A) / 2
	m := make(map[int]int)
	result := -1
	for _, i := range A {
		if count, ok := m[i]; ok {
			count = count + 1
			if count == n {
				result = i
				break
			}
			m[i] = count
		} else {
			m[i] = 1
		}
	}
	return result
}
