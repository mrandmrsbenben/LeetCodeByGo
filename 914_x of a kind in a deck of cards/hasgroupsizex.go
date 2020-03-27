package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards/
func main() {
	// common.MakeTestCases("hasGroupsSizeX")
	input1 := []int{1, 2, 3, 4, 4, 3, 2, 1}
	fmt.Printf("Expect1: true\n")
	fmt.Printf("Output1: %v\n", hasGroupsSizeX(input1))
	input2 := []int{1, 1, 1, 2, 2, 2, 3, 3}
	fmt.Printf("Expect2: false\n")
	fmt.Printf("Output2: %v\n", hasGroupsSizeX(input2))
	input3 := []int{1}
	fmt.Printf("Expect3: false\n")
	fmt.Printf("Output3: %v\n", hasGroupsSizeX(input3))
	input4 := []int{1, 1, 1, 1, 1}
	fmt.Printf("Expect4: true\n")
	fmt.Printf("Output4: %v\n", hasGroupsSizeX(input4))
	input5 := []int{1, 1, 2, 2, 2, 2}
	fmt.Printf("Expect5: true\n")
	fmt.Printf("Output5: %v\n", hasGroupsSizeX(input5))
	input6 := []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2}
	fmt.Printf("Expect6: true\n")
	fmt.Printf("Output6: %v\n", hasGroupsSizeX(input6))
	input7 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3}
	fmt.Printf("Expect7: true\n")
	fmt.Printf("Output7: %v\n", hasGroupsSizeX(input7))
}

//执行用时 :24 ms, 在所有 Go 提交中击败了93.75%的用户
//内存消耗 :5.9 MB, 在所有 Go 提交中击败了45.83%的用户
func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}
	sort.Ints(deck)
	if deck[0] == deck[len(deck)-1] {
		return true
	}
	var found bool
	for i := 2; i <= len(deck)/2; i++ {
		if len(deck)%i != 0 {
			continue
		} else if i > 2 && i%2 == 0 {
			continue
		}
		found = true
		for j := 0; j < len(deck); j += i {
			if deck[j] != deck[j+i-1] {
				found = false
				break
			}
		}
		if found {
			return true
		}
	}
	return false
}
