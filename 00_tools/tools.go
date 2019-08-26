package main

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
)

//[LeetCodeByGo: 更多LeetCode题库Go语言题解](https://github.com/mrandmrsbenben/LeetCodeByGo)
func main() {
	text, err := clipboard.ReadAll()
	if err == nil {
		s1 := strings.Split(text, "\n")
		fmt.Println(len(s1))
		s2 := "//"
		for i := 0; i < len(s1); i++ {
			// fmt.Printf(s1[i])
			if i == 5 {
				s2 = s2 + "\n"
				s2 = s2 + "//"
			}
			s2 += s1[i]
		}
		fmt.Printf(s2)
		clipboard.WriteAll(s2)
	}
}
