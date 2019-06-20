package main

import "fmt"

func main() {
	// common.MakeTestCases("isValid")
	input1 := "()"
	fmt.Printf("Expect1: true\n")
	fmt.Printf("Output1: %v\n", isValid(input1))
	input2 := "()[]{}"
	fmt.Printf("Expect2: true\n")
	fmt.Printf("Output2: %v\n", isValid(input2))
	input3 := "(])"
	fmt.Printf("Expect3: false\n")
	fmt.Printf("Output3: %v\n", isValid(input3))
	input4 := "([)]"
	fmt.Printf("Expect4: false\n")
	fmt.Printf("Output4: %v\n", isValid(input4))
	input5 := "{[]}"
	fmt.Printf("Expect5: true\n")
	fmt.Printf("Output5: %v\n", isValid(input5))
}

// 执行用时 : 4 ms, 在Valid Parentheses的Go提交中击败了21.73% 的用户
// 内存消耗 : 6.1 MB, 在Valid Parentheses的Go提交中击败了8.63% 的用户
func isValid(s string) bool {
	buf := ""
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			buf = buf + string(s[i])
		case ')':
			if len(buf) == 0 || buf[len(buf)-1] != '(' {
				return false
			}
			buf = buf[0 : len(buf)-1]
		case ']':
			if len(buf) == 0 || buf[len(buf)-1] != '[' {
				return false
			}
			buf = buf[0 : len(buf)-1]
		case '}':
			if len(buf) == 0 || buf[len(buf)-1] != '{' {
				return false
			}
			buf = buf[0 : len(buf)-1]
		}
	}
	return len(buf) == 0
}
