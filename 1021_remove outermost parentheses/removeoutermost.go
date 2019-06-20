package main

import "fmt"

type testcase struct {
	input  string
	output string
}

func main() {

	tests := []testcase{{"(()())(())", "()()()"},
		{"(()())(())(()(()))", "()()()()(())"},
		{"()()", ""}}
	for _, t := range tests {
		output := removeOuterParentheses(t.input)
		if output != t.output {
			fmt.Printf("Fail.Expect:%s, Actual:%s\n", t.output, output)
		} else {
			fmt.Printf("Passed.Output:%s\n", output)
		}
	}
}

func removeOuterParentheses(S string) string {
	in := []byte(S)
	output := ""
	prim := ""
	n := 0
	start := 0
	for i, c := range in {
		if c == '(' {
			n = n + 1
		} else {
			n = n - 1
		}
		if n == 0 {
			prim = S[start : i+1]
			output = output + prim[1:len(prim)-1]
			start = i + 1
		}
	}
	return output
}
