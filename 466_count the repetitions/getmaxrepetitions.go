package main

import "fmt"

func main() {
	// s1 := "acb"
	// n1 := 4
	// s2 := "ab"
	// n2 := 2
	s1 := "abaacdbac"
	n1 := 100
	s2 := "abcbd"
	n2 := 4
	// s1 := "aaa"
	// n1 := 3
	// s2 := "aa"
	// n2 := 1
	fmt.Println("Output: ", getMaxRepetitions(s1, n1, s2, n2))
}

//执行用时 :1140 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	c1, c2 := 0, 0
	j := 0
	for c1 < n1 {
		for i := range s1 {
			if s1[i] == s2[j] {
				if j == len(s2)-1 {
					j = 0
					c2++
				} else {
					j++
				}
			}
		}
		c1++
	}
	return c2 / n2
}
