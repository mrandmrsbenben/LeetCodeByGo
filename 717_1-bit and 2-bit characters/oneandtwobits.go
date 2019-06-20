package main

import "fmt"

func main() {
	bits := []int{0}
	fmt.Printf("Output: %v\n", isOneBitCharacter(bits))
}

func isOneBitCharacter(bits []int) bool {
	for i := 0; i < len(bits); {
		if bits[i] == 1 {
			if i == len(bits)-2 {
				return false
			}
			i += 2
		} else {
			if i == len(bits)-1 {
				return true
			}
			i++
		}
	}
	return false
}
