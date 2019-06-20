package main

import (
	"fmt"
)

func main() {
	fmt.Println("Output:", convertToTitle(1))
	fmt.Println("Output:", convertToTitle(28))
	fmt.Println("Output:", convertToTitle(701))
}

//执行用时 :0 ms, 在所有Go提交中击败了100.00%的用户
//内存消耗 :2 MB, 在所有Go提交中击败了56.25%的用户
func convertToTitle(n int) string {
	if n < 1 {
		return ""
	}
	letters := []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	s := ""
	if n%26 == 0 {
		s = letters[26] + s
		n = n - 26
	} else {
		s = letters[n%26] + s
		n = n - n%26
	}
	d := 26
	for {
		if n/d > 26 {
			s = letters[n/d%26] + s
			n -= d
		} else {
			s = letters[n/d] + s
			break
		}
		d *= 26
	}
	return s
}
