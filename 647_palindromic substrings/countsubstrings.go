package main

import "fmt"

func main() {
	s := "abccba"
	fmt.Println("Output: ", countSubstrings(s))
}

//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：2 MB, 在所有 Go 提交中击败了63.27%的用户
func countSubstrings(s string) int {
	cnt := len(s)
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt++
			cnt += countSubsPalindormic(i-2, i+1, s)
		}

		if i+1 < len(s) && s[i+1] == s[i-1] {
			cnt++
			cnt += countSubsPalindormic(i-2, i+2, s)
		}

	}

	return cnt
}

func countSubsPalindormic(j, k int, s string) int {
	cnt := 0
	for j >= 0 && k <= len(s)-1 && s[j] == s[k] {
		cnt++
		j--
		k++
	}
	return cnt
}
