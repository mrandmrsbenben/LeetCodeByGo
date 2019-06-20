package main

import "fmt"

func main() {
	s := "abcda"
	// s := "deeee"
	// s := "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"
	fmt.Println("Output:", validPalindrome(s))
}

//执行用时 :56 ms, 在所有 Go 提交中击败了47.54%的用户
//内存消耗 :7.5 MB, 在所有 Go 提交中击败了10.34%的用户
func validPalindrome(s string) bool {
	return isValid(s, false)
}

func isValid(s string, skipFlag bool) bool {
	i, j := 0, len(s)-1
	for i < len(s) && j >= 0 && i < j {
		if s[i] != s[j] {
			if skipFlag {
				return false
			}
			return isValid(s[0:i]+s[i+1:], true) || isValid(s[0:j]+s[j+1:], true)
		}
		i++
		j--
	}
	return true
}
