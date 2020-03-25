package main

import "strconv"

// https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/
func main() {
	findNthDigit(11)
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func findNthDigit(n int) int {
	str := "12345678910"
	if n <= 11 {
		d, _ := strconv.Atoi(str[n-1 : n])
		return d
	}
	length := 9
	i := 90
	j := 2
	for {
		if length+i*j > n {
			break
		}
		length += i * j
		i *= 10
		j++
	}
	d := pow10(j-1) + (n-length)/j
	x := (n - length) % j
	if x == 0 {
		return (d - 1) % 10
	}
	return d / pow10(j-x) % 10
}

func pow10(x int) int {
	n := 1
	for i := 0; i < x; i++ {
		n *= 10
	}
	return n
}
