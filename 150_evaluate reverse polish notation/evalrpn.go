package main

import (
	"fmt"
	"strconv"
)

func main() {
	// tokens := []string{"2", "1", "+", "3", "*"}
	// tokens := []string{"4", "13", "5", "/", "+"}
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}

//执行用时 :8 ms, 在所有 Go 提交中击败了75.89%的用户
//内存消耗 :4.2 MB, 在所有 Go 提交中击败了58.49%的用户
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	var n, x, y int
	for _, t := range tokens {
		if len(stack) > 1 {
			x, y = stack[len(stack)-2], stack[len(stack)-1]
		}
		switch t {
		case "+":
			n = x + y
			stack = pop2(stack)
		case "-":
			n = x - y
			stack = pop2(stack)
		case "*":
			n = x * y
			stack = pop2(stack)
		case "/":
			n = x / y
			stack = pop2(stack)
		default:
			n, _ = strconv.Atoi(t)
		}
		stack = append(stack, n)
	}
	return stack[0]
}

func pop2(stack []int) []int {
	if len(stack) <= 2 {
		return []int{}
	}
	return stack[0 : len(stack)-2]
}
