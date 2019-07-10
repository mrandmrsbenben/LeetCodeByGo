package main

import "fmt"

func main() {
	fmt.Println("Output:", partition("aabb"))
}

//执行用时 :40 ms, 在所有 Go 提交中击败了49.28%的用户
//内存消耗 :7.3 MB, 在所有 Go 提交中击败了71.93%的用户
// 思路：从头至尾逐位分割字符串，分割后前半部分为回文串时递归调用函数，找出后半部分所有的回文串组合。
func partition(s string) [][]string {
	if len(s) == 1 {
		return [][]string{{s}}
	}
	ps := make([][]string, 0)
	var subps [][]string
	for i := range s {
		if isPalindrome(s[0 : i+1]) {
			if i == len(s)-1 {
				ps = append(ps, []string{s[0 : i+1]})
			} else {
				subps = partition(s[i+1:])
				for j := range subps {
					ps = append(ps, append([]string{s[0 : i+1]}, subps[j]...))
				}
			}
		}
	}
	return ps
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
