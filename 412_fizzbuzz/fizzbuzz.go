package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 15
	fmt.Printf("Output: %v\n", fizzBuzz(n))
}

func fizzBuzz(n int) []string {
	s := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			s[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			s[i-1] = "Fizz"
		} else if i%5 == 0 {
			s[i-1] = "Buzz"
		} else {
			s[i-1] = strconv.Itoa(i)
		}
	}
	return s
}
