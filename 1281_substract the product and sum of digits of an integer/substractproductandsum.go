package main

import "fmt"

func main() {
	n := 4421
	fmt.Println("Output: ", subtractProductAndSum(n))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了20.74%的用户
func subtractProductAndSum(n int) int {
	pro, sum := 1, 0
	var x int
	for n > 0 {
		x = n % 10
		sum += x
		pro *= x
		n /= 10
	}
	return pro - sum
}
