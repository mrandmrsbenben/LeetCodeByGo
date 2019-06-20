package main

import "fmt"

func main() {
	N := 2
	fmt.Printf("Output: %v\n", divisorGame(N))
}

func divisorGame(N int) bool {
	if N%2 != 0 {
		return false
	}
	return true
}
