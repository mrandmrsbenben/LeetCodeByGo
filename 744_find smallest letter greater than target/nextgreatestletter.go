package main

import "fmt"

func main() {
	letters := []byte{'c', 'f', 'j'}
	target := byte('k')
	fmt.Printf("Output: %c\n", nextGreatestLetter(letters, target))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	index := -1
	for i := range letters {
		if letters[i] > target {
			index = i
			break
		}
	}
	if index == -1 {
		index = 0
	}
	return letters[index]
}
