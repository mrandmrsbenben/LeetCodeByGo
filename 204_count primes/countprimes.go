package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println("Output:", countPrimes(499739))
}

//执行用时 :12 ms, 在所有 Go 提交中击败了94.70%的用户
//内存消耗 :5 MB, 在所有 Go 提交中击败了50.91%的用户
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	cnt := n - 2
	nums := make([]bool, n)
	m := int(math.Sqrt(float64(n)))
	for i := 2; i <= m; i++ {
		if !nums[i] {
			for j := i; i*j < n; j++ {
				if !nums[i*j] {
					nums[i*j] = true
					cnt--
				}
			}
		}
	}
	return cnt
}

//执行用时 :64 ms, 在所有 Go 提交中击败了47.68%的用户
//内存消耗 :26.4 MB, 在所有 Go 提交中击败了5.45%的用户
func countPrimesV1(n int) int {
	if n < 3 {
		return 0
	}
	cnt := n - 2
	nums := make([]int, n)
	m := int(math.Sqrt(float64(n)))
	for i := 2; i <= m; i++ {
		if nums[i] == 0 {
			for j := i; i*j < n; j++ {
				if nums[i*j] == 0 {
					nums[i*j] = 1
					cnt--
				}
			}
		}
	}
	return cnt
}

// Time Limit Exceeded
func countPrimesSlow(n int) int {
	if n == 1 {
		return 0
	}
	primes := make([]int, n-1)
	copy(primes, []int{2, 3, 5, 7, 11})
	if n < 12 {
		return sort.SearchInts(primes, n)
	}
	var isPrime bool
	index := 5
	for i := 12; i < n; i++ {
		isPrime = true
		for j := range primes {
			if primes[j] == 0 {
				break
			} else if i%primes[j] == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes[index] = i
			index++
		}
	}
	return index
}
