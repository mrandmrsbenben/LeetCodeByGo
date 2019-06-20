package main

import "fmt"

func main() {
	m := 3
	n := 3
	ops := [][]int{}
	fmt.Printf("Output: %d\n", maxCount(m, n, ops))
}

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	// s := make([][]int, m)
	// for i := range s {
	// 	s[i] = make([]int, n)
	// }
	// for _, op := range ops {
	// 	for i := 0; i < op[0]; i++ {
	// 		for j := 0; j < op[1]; j++ {
	// 			s[i][j] = s[i][j] + 1
	// 		}
	// 	}
	// 	fmt.Println(s)
	// }
	mina, minb := ops[0][0], ops[0][1]
	for _, op := range ops {
		if mina > op[0] {
			mina = op[0]
		}
		if minb > op[1] {
			minb = op[1]
		}
	}
	return mina * minb
}
