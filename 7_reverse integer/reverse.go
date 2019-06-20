package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	testcase := []int{123, -123, 120, 7463847412}
	for i, t := range testcase {
		fmt.Println(i, ":", reverse(t))
	}
}

func reverse(x int) int {
	r := 0
	str := strconv.Itoa(x)
	l := len(str)
	ch := []byte(str)
	revstr := make([]byte, l)
	start := 0
	if ch[0] == '-' {
		revstr[0] = ch[0]
		start = 1
	}
	for i := start; i < l; i++ {
		revstr[l-i-1+start] = ch[i]
	}
	r, err := strconv.Atoi(string(revstr))
	if err != nil || r > math.MaxInt32 || r < math.MinInt32 {
		r = 0
	}
	return r
}
