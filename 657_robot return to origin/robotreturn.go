package main

import (
	"fmt"
	"strings"
)

type testcase struct {
	moves  string
	origin bool
}

func main() {
	cases := []testcase{{"UD", true}, {"RLR", false}, {"LDUR", true}}
	var origin bool
	for _, t := range cases {
		origin = judgeCircle(t.moves)
		if origin != t.origin {
			fmt.Printf("Fail.Moves:%s, Expected:%v, Actual:%v\n",
				t.moves, t.origin, origin)
		} else {
			fmt.Printf("Pass.Moves:%s, Origin:%v\n", t.moves, origin)
		}
	}
}

func judgeCircle(moves string) bool {
	x := strings.Count(moves, "L") - strings.Count(moves, "R")
	y := strings.Count(moves, "U") - strings.Count(moves, "D")
	if x == 0 && y == 0 {
		return true
	}
	return false
}
