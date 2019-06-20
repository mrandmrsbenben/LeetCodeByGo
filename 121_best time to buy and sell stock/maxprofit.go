package main

import (
	"fmt"
	"sort"
)

func main() {
	// common.MakeTestCases("maxProfit")
	input1 := []int{6, 9, 5, 10, 2, 4}
	fmt.Printf("Expect1: 5\n")
	fmt.Printf("Output1: %v\n", maxProfit(input1))
	input2 := []int{7, 6, 4, 3, 1}
	fmt.Printf("Expect2: 0\n")
	fmt.Printf("Output2: %v\n", maxProfit(input2))
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	max := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}

func maxProfit0(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := 0
	var p []int
	for i := range prices {
		p = make([]int, len(prices)-i)
		copy(p, prices[i:])
		sort.Ints(p)
		if max < p[len(p)-1]-prices[i] {
			max = p[len(p)-1] - prices[i]
		}
	}
	return max
}
