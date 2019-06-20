package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	N := 20
	fmt.Printf("Output: %d\n", rotatedDigits(N))
}

func rotatedDigits(N int) int {
	cnt := 0
	for i := 1; i <= N; i++ {
		s := strconv.Itoa(i)
		if strings.ContainsAny(s, "347") || !strings.ContainsAny(s, "2569") {
			continue
		}
		fmt.Println(i)
		cnt = cnt + 1
	}
	return cnt
}
