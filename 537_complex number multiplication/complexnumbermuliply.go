package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// a, b := "1+1i", "1+1i"
	a, b := "1+-1i", "1+-1i"
	fmt.Println("Output:", complexNumberMultiply(a, b))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func complexNumberMultiply(a string, b string) string {
	x := strToComplex(a)
	y := strToComplex(b)
	c1 := x[0]*y[0] - x[1]*y[1]
	c2 := x[0]*y[1] + x[1]*y[0]
	return strconv.Itoa(c1) + "+" + strconv.Itoa(c2) + "i"
}

func strToComplex(a string) [2]int {
	comp := [2]int{}
	s := strings.Split(a, "+")
	comp[0], _ = strconv.Atoi(s[0])
	comp[1], _ = strconv.Atoi(strings.TrimSuffix(s[1], "i"))
	return comp
}
