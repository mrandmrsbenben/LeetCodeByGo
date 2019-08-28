package main

import "fmt"

func main() {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(T))
}

func dailyTemperatures(T []int) []int {
	l := len(T)
	days := make([]int, l)
	days[l-1] = 0
	stack := []int{l - 1}
	var index int
	for i := l - 2; i >= 0; i-- {
		index = -1
		for j := len(stack) - 1; j >= 0; j-- {
			if T[i] < T[stack[j]] {
				index = j
				break
			}
		}
		if index == -1 {
			stack = []int{}
		} else {
			stack = stack[0 : index+1]
		}
		if len(stack) == 0 {
			days[i] = 0
		} else {
			days[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return days
}

//执行用时 :2964 ms, 在所有 Go 提交中击败了18.00%的用户
//内存消耗 :13.8 MB, 在所有 Go 提交中击败了86.36%的用户
func dailyTemperaturesV1(T []int) []int {
	days := make([]int, len(T))
	var cnt int
	for i := range T {
		cnt = 0
		for j := i + 1; j < len(T); j++ {
			cnt++
			if T[j] > T[i] {
				days[i] = cnt
				break
			}
		}
	}
	return days
}
