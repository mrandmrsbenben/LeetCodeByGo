package main

import "fmt"

func main() {
	// common.MakeTestCases("lemonadeChange")
	input1 := []int{5, 5, 5, 10, 20}
	fmt.Printf("Expect1: true\n")
	fmt.Printf("Output1: %v\n", lemonadeChange(input1))
	input2 := []int{5, 5, 10}
	fmt.Printf("Expect2: true\n")
	fmt.Printf("Output2: %v\n", lemonadeChange(input2))
	input3 := []int{10, 10}
	fmt.Printf("Expect3: false\n")
	fmt.Printf("Output3: %v\n", lemonadeChange(input3))
	input4 := []int{5, 5, 10, 10, 20}
	fmt.Printf("Expect4: false\n")
	fmt.Printf("Output4: %v\n", lemonadeChange(input4))
}

func lemonadeChange(bills []int) bool {
	c5, c10 := 0, 0
	for _, b := range bills {
		switch b {
		case 5:
			c5 = c5 + 1
		case 10:
			if c5 == 0 {
				return false
			}
			c5 = c5 - 1
			c10 = c10 + 1
		case 20:
			if c10 > 0 && c5 > 0 {
				c10 = c10 - 1
				c5 = c5 - 1
			} else if c5 >= 3 {
				c5 = c5 - 3
			} else {
				return false
			}
		}
	}
	return true
}
