package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	left := 9999
	right := 10000
	fmt.Printf("Self Dividing Numbers:%v\n", selfDividingNumbers(left, right))
}

func selfDividingNumbers(left int, right int) []int {
	sdn := make([]int, 0)
	var ns string
	var n int
	var canSelfDiv bool
	for i := left; i <= right; i++ {
		ns = strconv.Itoa(i)
		if strings.Contains(ns, "0") {
			continue
		}
		canSelfDiv = true
		for j := 0; j < len(ns); j++ {
			n, _ = strconv.Atoi(string(ns[j]))
			if i%n != 0 {
				canSelfDiv = false
				break
			}
		}
		if canSelfDiv {
			sdn = append(sdn, i)
		}
	}
	return sdn
}
