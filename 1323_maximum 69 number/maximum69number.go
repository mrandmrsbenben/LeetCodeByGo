package main

import "fmt"

func main() {
	num := 9999
	fmt.Println("Output: ", maximum69Number(num))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func maximum69Number(num int) int {
	n := 1000
	m := 0
	for n > 0 {
		if num/n == 9 {
			m += 9 * n
			num = num % n
		} else if num/n == 6 {
			m += 9*n + num%n
			break
		}
		n /= 10
	}
	return m
}
