package main

import "fmt"

func main() {
	fmt.Println("Output:", tribonacci(4))
	fmt.Println("Output:", tribonacci(25))
}

func tribonacci(n int) int {
	nm := make(map[int]int)
	var trib func(n int) int
	trib = func(n int) int {
		switch n {
		case 0:
			return 0
		case 1, 2:
			return 1
		default:
			if v, ok := nm[n]; ok {
				return v
			}
			v := trib(n-1) + trib(n-2) + trib(n-3)
			nm[n] = v
			return v
		}
	}
	return trib(n)
}

func tribonacciV1(n int) int {
	ns := make([]int, n+1)
	ns[0] = 0
	if n >= 2 {
		ns[1], ns[2] = 1, 1
	} else if n >= 1 {
		ns[1] = 1
	}
	for i := 3; i <= n; i++ {
		ns[i] = ns[i-3] + ns[i-2] + ns[i-1]
	}
	return ns[n]
}
