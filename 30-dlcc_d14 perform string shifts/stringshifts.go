package main

import "fmt"

func main() {
	s := "abcdefg"
	shift := [][]int{{1, 1}, {1, 1}, {0, 2}, {0, 3}}
	fmt.Println("Output: ", stringShift(s, shift))
}

func stringShift(s string, shift [][]int) string {
	res := s
	cnt := 0
	for i := range shift {
		if shift[i][0] == 0 {
			cnt -= shift[i][1]
		} else {
			cnt += shift[i][1]
		}
	}
	cnt %= len(s)
	if cnt > 0 {
		// shift to right
		res = s[len(s)-cnt:] + res[0:len(s)-cnt]
	} else if cnt < 0 {
		// shift to left
		cnt *= -1
		res = res[cnt:] + s[0:cnt]
	}
	return res
}
