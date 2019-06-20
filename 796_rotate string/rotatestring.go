package main

import "fmt"

func main() {
	A := ""
	B := ""
	fmt.Printf("Output: %v\n", rotateString(A, B))
}

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	} else if len(A) == 0 {
		return true
	}
	for i := 0; i < len(A); i++ {
		if A[i+1:]+A[0:i+1] == B {
			return true
		}
	}
	return false
}
