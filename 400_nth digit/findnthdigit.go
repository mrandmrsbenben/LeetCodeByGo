package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println("Output:", findNthDigit(11))
	// fmt.Println("Output:", findNthDigit(12))
	// fmt.Println("Output:", findNthDigit(13))
	// fmt.Println("Output:", findNthDigit(30))
	// fmt.Println("Output:", findNthDigit(100))
	// fmt.Println("Output:", findNthDigit(101))
	// fmt.Println("Output:", findNthDigit(102))
	// fmt.Println("Output:", findNthDigit(103))
	// fmt.Println("Output:", findNthDigit(1000))
	fmt.Println("Output:", findNthDigit(999999999))
	// fmt.Println("Output:", findNthDigit(10000000))
}

//执行用时 :0 ms, 在所有Go提交中击败了100.00%的用户
//内存消耗 :2 MB, 在所有Go提交中击败了16.67%的用户
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

// Too Slow Time limit exceeded
func findNthDigitSlow(n int) int {
	str := "012345678910"
	if n >= 12 {
		for i := 11; i < n && len(str) <= n; i++ {
			str += strconv.Itoa(i)
		}
	}
	fmt.Println(str)
	d, _ := strconv.Atoi(str[n : n+1])
	return d
}
