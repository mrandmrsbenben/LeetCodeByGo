package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	L := 990
	R := 1048
	fmt.Printf("Output: %d\n", countPrimeSetBits(L, R))
}
func countPrimeSetBits(L int, R int) int {
	m := map[int]int{2: 0, 3: 0, 5: 0, 7: 0, 11: 0, 13: 0, 17: 0, 19: 0}
	count := R - L + 1
	var n int
	for i := L; i <= R; i++ {
		n = strings.Count(fmt.Sprintf("%b", i), "1")
		if _, ok := m[n]; !ok {
			count = count - 1
		}
	}
	return count
}
func countPrimeSetBits0(L int, R int) int {
	count := R - L + 1
	var n int
	// primes := make(map[int]bool)
	// primes[2] = true
	for i := L; i <= R; i++ {
		n = strings.Count(fmt.Sprintf("%b", i), "1")
		if n == 2 {
			continue
		} else if n == 1 || n%2 == 0 {
			count = count - 1
			continue
		}
		// v, ok := primes[n]
		// if ok {
		// 	if !v {
		// 		count = count - 1
		// 	}
		// 	continue
		// }
		// primes[n] = true
		for j := 2; j <= int(math.Ceil(math.Sqrt(float64(n)))); j++ {
			if n%j == 0 {
				count = count - 1
				// primes[n] = false
				break
			}
		}
	}
	return count
}
