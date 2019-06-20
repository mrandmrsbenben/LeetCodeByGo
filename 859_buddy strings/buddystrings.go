package main

import "fmt"

func main() {
	// A := "aaaaaaabc"
	// B := "aaaaaaacb"
	// A := "ab"
	// B := "ba"
	A := "abcd"
	B := "abcd"
	// A := ""
	// B := "ab"
	fmt.Println("Output:", buddyStrings(A, B))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.7 MB, 在所有 Go 提交中击败了21.43%的用户
func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	} else if len(A) < 2 {
		return false
	}
	var str string
	for i := 0; i < len(A); i++ {
		if i > 0 && A[i] == A[i-1] {
			continue
		}
		for j := i + 1; j < len(A); j++ {
			if j > i+1 && A[j] == A[j-1] {
				continue
			}
			str = A[0:i] + A[j:j+1] + A[i+1:j] + A[i:i+1] + A[j+1:]
			if str == B {
				return true
			}
		}
	}
	return false
}
