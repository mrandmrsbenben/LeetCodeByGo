package main

import "fmt"

func main() {
	fmt.Printf("F(%d) = %d\n", 30, fib(30))
}

func fibN(N int) int {
	if N == 0 {
		return 0
	}
	f := make([]int, N+1)
	f[0], f[1] = 0, 1
	for i := 2; i <= N; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[N]
}

func fib(N int) int {
	if N == 0 || N == 1 {
		return N
	}
	return fib(N-1) + fib(N-2)
}
