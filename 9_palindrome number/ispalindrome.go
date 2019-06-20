package main

import (
	"fmt"
	"strconv"
)

func main() {
	// common.MakeTestCases("isPalindrome")
	input1 := 121
	fmt.Printf("Expect1: true\n")
	fmt.Printf("Output1: %v\n", isPalindrome(input1))
	input2 := -121
	fmt.Printf("Expect2: false\n")
	fmt.Printf("Output2: %v\n", isPalindrome(input2))
	input3 := 10
	fmt.Printf("Expect3: false\n")
	fmt.Printf("Output3: %v\n", isPalindrome(input3))
}

func isPalindromeS(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	n := make([]int, 0)
	n = append(n, x%10)
	for x >= 10 {
		x = x / 10
		n = append(n, x%10)
	}
	for i := 0; i < len(n)/2; i++ {
		if n[i] != n[len(n)-1-i] {
			return false
		}
	}
	return true
}
