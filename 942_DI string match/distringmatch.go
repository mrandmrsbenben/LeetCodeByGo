package main

import "fmt"

type testcase struct {
	distr  string
	expect []int
}

func main() {
	cases := []testcase{{"IDID", []int{0, 4, 1, 3, 2}},
		{"III", []int{0, 1, 2, 3}},
		{"DDI", []int{3, 2, 0, 1}}}
	for _, t := range cases {
		fmt.Printf("Di String:%s, Output Expect:%v, Actual:%v\n",
			t.distr, t.expect, diStringMatch(t.distr))
	}
}

func diStringMatch(S string) []int {
	dc := len(S)
	ic := 0
	a := make([]int, len(S)+1)
	for i := 0; i < len(S); i++ {
		ch := string(S[i])
		switch ch {
		case "D":
			a[i] = dc
			dc = dc - 1
		case "I":
			a[i] = ic
			ic = ic + 1
		}
	}
	a[len(S)] = dc
	return a
}
