package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := "1234567890"
	k := 9
	// num := "1432219"
	// k := 3
	// num := "10200"
	// num := "112"
	// k := 1
	fmt.Println("Output: ", removeKdigits(num, k))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :3.1 MB, 在所有 Go 提交中击败了100.00%的用户
func removeKdigits(num string, k int) string {
	res := make([]rune, 0)

	for _, c := range num {
		for k > 0 && len(res) > 0 && res[len(res)-1] > c {
			res = res[:len(res)-1]
			k--
		}
		res = append(res, c)
	}
	res = res[:len(res)-k]
	s := strings.TrimLeft(string(res), "0")
	if s == "" {
		s = "0"
	}
	return s
}

func removeKdigitsV0(num string, k int) string {
	if len(num) == k {
		return "0"
	}

	minS := make(map[string]string)
	var remove func(string, int) string
	remove = func(n string, k int) string {
		key := n + "," + strconv.Itoa(k)
		if v, ok := minS[key]; ok {
			return v
		}

		res := n
		if k >= len(n) {
			res = ""
		} else if k > 0 && len(n) > 1 {
			res = min(remove(n[1:], k-1), n[0:1]+remove(n[1:], k))
		}
		minS[key] = res
		return res
	}
	res := remove(num, k)

	index := -1
	for i := 0; i < len(res)-1; i++ {
		if res[i] != '0' {
			break
		}
		index = i
	}
	if index != -1 {
		res = res[index+1:]
	}
	return res
}

func min(a, b string) string {
	if a != b {
		for i := range a {
			if a[i] < b[i] {
				return a
			} else if a[i] > b[i] {
				return b
			}
		}
	}
	return a
}
