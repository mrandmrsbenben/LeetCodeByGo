package main

import "fmt"

func main() {
	// s, t := "ADOBECOEBANC", "ABC"
	s, t := "abc", "ac"
	fmt.Println("Output: ", minWindow(s, t))
}

//执行用时 :32 ms, 在所有 Go 提交中击败了41.48%的用户
//内存消耗 :6.3 MB, 在所有 Go 提交中击败了50.00%的用户
func minWindow(s string, t string) string {
	if s == t {
		return s
	}

	res := ""
	ms, mt := make(map[byte]int), make(map[byte]int)
	for i := range t {
		mt[t[i]]++
	}

	cnt, min := 0, len(s)+1
	ps := make([]int, 0)
	l, j := -1, 0
	for i := range s {
		if mt[s[i]] > 0 {
			if ms[s[i]] < mt[s[i]] {
				cnt++
			}
			ms[s[i]]++
			ps = append(ps, i)
		}
		if cnt == len(t) {
			if l == -1 {
				l = ps[0]
			}
			for {
				ms[s[l]]--
				if i-l+1 < min {
					min = i - l + 1
					res = s[l : i+1]
				}
				if ms[s[l]] < mt[s[l]] {
					break
				}
				j++
				l = ps[j]
			}
			if len(res) == 1 {
				break
			}
			j++
			l = ps[j]
			cnt--
		}
	}
	return res
}
