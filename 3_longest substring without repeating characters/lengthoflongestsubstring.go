package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("Output:", lengthOfLongestSubstring("abcabcbb"))
	// fmt.Println("Output:", lengthOfLongestSubstring("jxdlnaaij"))
	// fmt.Println("Output:", lengthOfLongestSubstring("pwwkew"))
	// fmt.Println("Output:", lengthOfLongestSubstring("cdd"))
	// fmt.Println("Output:", lengthOfLongestSubstring("ohomm"))
	// fmt.Println("Output:", lengthOfLongestSubstring("ddddd"))
	fmt.Println("Output:", lengthOfLongestSubstring(" "))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxLen := 1
	foundChars := make(map[byte]int)
	i, j := 0, 0
	for j = range s {
		currChar := s[j]
		if lastIndex, ok := foundChars[currChar]; ok {
			if j-i > maxLen {
				maxLen = j - i
			}
			if lastIndex+1 > i {
				i = lastIndex + 1
			}
		}
		foundChars[currChar] = j
	}
	if j-i+1 > maxLen {
		maxLen = j - i + 1
	}
	return maxLen
}

func lengthOfLongestSubstringV1(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxLen := 1
	l := len(s)
	var substr string
	var index int
	for i := 0; i < l-1 && maxLen < l-i; i++ {
		index = strings.Index(s[i+1:], string(s[i]))
		if index == -1 {
			if isNoRepeat(s[i:]) {
				if l-i > maxLen {
					maxLen = l - i
				}
			} else {
				substr = s[i:]
				for len(substr) > maxLen {
					if isNoRepeat(substr) {
						if len(substr) > maxLen {
							maxLen = len(substr)
						}
						break
					}
					substr = substr[0 : len(substr)-1]
				}
			}
		} else if index >= 0 {
			substr = s[i : i+1+index]
			if isNoRepeat(substr) {
				if len(substr) > maxLen {
					maxLen = len(substr)
				}
			} else {
				substr = s[i:]
				for len(substr) > maxLen {
					if isNoRepeat(substr) {
						if len(substr) > maxLen {
							maxLen = len(substr)
						}
						break
					}
					substr = substr[0 : len(substr)-1]
				}
			}
		}
	}
	return maxLen
}

func isNoRepeat(s string) bool {
	var c string
	for i := range s {
		c = string(s[i])
		if strings.Index(s, c) != strings.LastIndex(s, c) {
			return false
		}
	}
	return true
}
