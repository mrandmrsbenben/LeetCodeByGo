package main

import (
	"fmt"
	"sort"
)

func main() {
	// s := "catsanddog"
	// wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Println("Output: ", wordBreak(s, wordDict))
}

//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：2.6 MB, 在所有 Go 提交中击败了100.00%的用户
func wordBreak(s string, wordDict []string) []string {
	if len(s) == 0 || wordDict == nil || len(wordDict) == 0 {
		return []string{}
	}

	sort.Strings(wordDict)
	wm := make(map[string][]string)

	var wb func(string) []string
	wb = func(str string) []string {
		if ret, ok := wm[str]; ok {
			return ret
		}

		ret := []string{}
		var strs []string
		for _, w := range wordDict {
			if str == w {
				ret = append(ret, w)
				break
			} else if len(str) >= len(w) && str[0:len(w)] == w {
				strs = wb(str[len(w):])
				for i := range strs {
					ret = append(ret, w+" "+strs[i])
				}
			}
		}
		wm[str] = ret
		return ret
	}

	return wb(s)
}
