package main

import "fmt"

func main() {
	fmt.Println("Output: ", myPow(2.0000, 10))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户
func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPow(x, n*-1)
	}
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	h := pow(x, n/2)
	if n%2 == 0 {
		return h * h
	}
	return x * h * h
}
