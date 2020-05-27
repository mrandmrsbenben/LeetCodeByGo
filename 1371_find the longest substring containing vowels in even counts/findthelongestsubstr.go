package main

import "fmt"

func main() {
	// s := "eleetminicoworoep"
	// s := "leetcodeisgreat"
	// s := "bcbcbc"
	// s := "aeioe"
	s := "qnplnlarrtztkotazhufrsfczrzibvccaoayyihidztfljcffiqfviuwjowkppdajmknzgidixqgtnahamebxfowqvnrhuzwqohquamvszkvunbxjegbjccjjxfnsiearbsgsofywtqbmgldgsvnsgpdvmjqpaktmjafgkzszekngivdmrlvrpyrhcxbceffrgiyktqilkkdjhtywpesrydkbncmzeekdtszmcsrhsciljsrdoidzbjatvacndzbghzsnfdofvhfxdnmzrjriwpkdgukbaazjxtkomkmccktodig"
	fmt.Println("Output: ", findTheLongestSubstring(s))
}

//执行用时 :28 ms, 在所有 Go 提交中击败了30.57%的用户
//内存消耗 :6.4 MB, 在所有 Go 提交中击败了100.00%的用户
func findTheLongestSubstring(s string) int {

	dic := make([][]int, 5)
	for i := range dic {
		dic[i] = make([]int, 0)
	}
	for i := range s {
		switch s[i] {
		case 'a':
			dic[0] = append(dic[0], i)
		case 'e':
			dic[1] = append(dic[1], i)
		case 'i':
			dic[2] = append(dic[2], i)
		case 'o':
			dic[3] = append(dic[3], i)
		case 'u':
			dic[4] = append(dic[4], i)
		}
	}

	buf := make(map[string]string)
	maxLen := 0
	var findLongest func(int, int) string
	findLongest = func(l, r int) string {
		if l > r {
			return ""
		} else if l == r {
			if s[l] == 'a' || s[l] == 'e' || s[l] == 'i' || s[l] == 'o' || s[l] == 'u' {
				return ""
			}
			return s[l : r+1]
		} else if r-l+1 <= maxLen {
			return ""
		}

		key, ans := s[l:r+1], s[l:r+1]
		if _, ok := buf[ans]; ok {
			return buf[ans]
		}
		for i := range dic {
			cnt := 0
			first, last := -1, -1
			for j := range dic[i] {
				if dic[i][j] >= l && dic[i][j] <= r {
					if first == -1 {
						first = dic[i][j]
					}
					last = dic[i][j]
					cnt++
				}
			}
			if cnt%2 == 1 {
				if first == last {
					ans = max(findLongest(0, first-1), findLongest(last+1, r))
				} else {
					ans = max(max(findLongest(0, first-1), findLongest(first+1, r)),
						max(findLongest(0, last-1), findLongest(last+1, r)))
				}
			}
		}
		buf[key] = ans
		if len(ans) > maxLen {
			maxLen = len(ans)
		}
		return ans
	}
	findLongest(0, len(s)-1)
	return maxLen
}

func max(a, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}
