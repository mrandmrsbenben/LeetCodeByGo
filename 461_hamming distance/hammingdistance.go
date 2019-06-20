package main

import (
	"fmt"
	"strings"
)

type testcase struct {
	x, y   int
	expect int
}

func main() {
	cases := []testcase{{1, 4, 2}, {2, 4, 2}}
	for _, t := range cases {
		fmt.Printf("Input: %d, %d Output expect:%d actual:%d\n",
			t.x, t.y, t.expect, hammingDistance(t.x, t.y))
	}
}

func hammingDistance(x int, y int) int {
	return strings.Count(fmt.Sprintf("%b", x^y), "1")
}
