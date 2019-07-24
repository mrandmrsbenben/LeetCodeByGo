package main

import "fmt"

func main() {
	fmt.Println("Output:", myPow(2, 10))
	fmt.Println("Output:", myPow(2.1, 3))
	fmt.Println("Output:", myPow(2, -2))
	fmt.Println("Output:", myPow(0.00001, 2147483647))
}

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
