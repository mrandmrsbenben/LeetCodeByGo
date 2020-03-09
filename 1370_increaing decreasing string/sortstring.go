package main

import "fmt"

func main() {
	fmt.Println("Output: ", sortString("ops"))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :6.3 MB, 在所有 Go 提交中击败了100.00%的用户
func sortString(s string) string {
	abs := [26]int{}
	res := ""
	for _, c := range s {
		abs[c-'a']++
	}

	l := len(s)
	dec := false
	for l > 0 {
		if !dec {
			for i := 0; i < 26 && l > 0; i++ {
				if abs[i] > 0 {
					res += string('a' + i)
					abs[i]--
					l--
				}
			}
		} else {
			for i := 25; i >= 0 && l > 0; i-- {
				if abs[i] > 0 {
					res += string('a' + i)
					abs[i]--
					l--
				}
			}
		}
		dec = !dec
	}
	return res
}
