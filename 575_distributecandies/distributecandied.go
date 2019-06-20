package main

import "fmt"

func main() {
	// common.MakeTestCases("distributeCandies")
	candies1 := []int{1, 1, 2, 2, 3, 3}
	fmt.Printf("Output: %v\n", distributeCandies(candies1))
	fmt.Printf("Expect: 3\n")
	candies2 := []int{1, 1, 2, 3}
	fmt.Printf("Output: %v\n", distributeCandies(candies2))
	fmt.Printf("Expect: 2\n")
}

func distributeCandies(candies []int) int {
	m := make(map[int]int)
	for _, v := range candies {
		if _, ok := m[v]; !ok {
			m[v] = v
		}
	}
	if len(m) >= len(candies)/2 {
		return len(candies) / 2
	}
	return len(m)
}
