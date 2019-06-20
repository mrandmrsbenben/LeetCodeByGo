package main

import "fmt"

func main() {
	S := "dcabbacfg"
	fmt.Printf("Output: %v\n", removeDuplicates(S))
}

func removeDuplicates(S string) string {
	flag := false
	for i := 0; i < len(S) && !flag; i++ {
		if i == len(S)-1 {
			flag = false
		} else {
			if S[i:i+1] == S[i+1:i+2] {
				for j := 0; i-j >= 0 && i+j < len(S) && !flag; j++ {
					if S[i-j] == S[i+j+1] {
						if i-j == 0 {
							S = S[i+1+j+1:]
							flag = true
						} else if i+j+1 == len(S)-1 {
							S = S[0 : i-j]
							flag = true
						}
					} else {
						if i-j == 0 {
							S = S[0:1] + S[i+j+1:]
						} else {
							S = S[0:i-j+1] + S[i+j+1:]
						}
						flag = true
					}
				}
			}
		}
	}
	if flag {
		return removeDuplicates(S)
	}
	return S
}

func removeDuplicatesSlow(S string) string {
	newS := ""
	flag := false
	for i := 0; i < len(S); i++ {
		if i == len(S)-1 {
			newS = newS + S[i:]
		} else {
			if S[i:i+1] == S[i+1:i+2] {
				i = i + 1
				flag = true
			} else {
				newS = newS + S[i:i+1]
			}
		}
	}
	if flag {
		return removeDuplicates(newS)
	}
	return newS
}
