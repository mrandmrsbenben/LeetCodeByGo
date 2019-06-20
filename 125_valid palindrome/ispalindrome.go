package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// s := "!A man, a plan, a canal: Panama!"
	// s := "0P"
	s := "rac a car"
	fmt.Printf("Output: %v\n", isPalindrome(s))
}

// 执行用时 : 4 ms, 在Valid Palindrome的Go提交中击败了97.55% 的用户
// 内存消耗 : 5.4 MB, 在Valid Palindrome的Go提交中击败了20.00% 的用户
func isPalindrome(s string) bool {
	a := strings.FieldsFunc(strings.ToLower(s), func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})
	s2 := strings.Join(a, "")
	for i := range s2 {
		if s2[i] != s2[len(s2)-1-i] {
			return false
		}
	}
	return true
}
func isPalindromeV1(s string) bool {
	// s = strings.ToLower(s)
	s2 := ""
	for i := range s {
		fmt.Println(s[i], string(s[i]))
		if unicode.IsLetter(rune(s[i])) || unicode.IsNumber(rune(s[i])) {
			s2 += string(unicode.ToLower(rune(s[i])))
		}
	}
	for i := range s2 {
		if s2[i] != s2[len(s2)-1-i] {
			return false
		}
	}
	return true
}
