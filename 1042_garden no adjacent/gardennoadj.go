package main

import "fmt"

func main() {
	// N := 3
	// paths := [][]int{{1, 2}, {2, 3}, {3, 1}}
	// N := 4
	// paths := [][]int{{1, 2}, {3, 4}}
	// N := 4
	// paths := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}, {2, 4}}
	N := 1
	paths := [][]int{}
	fmt.Printf("Output: %v\n", gardenNoAdj(N, paths))
}

// 执行用时 :396 ms, 在所有Go提交中击败了57.89%的用户
// 内存消耗 :13.5 MB, 在所有Go提交中击败了100.00%的用户
func gardenNoAdj(N int, paths [][]int) []int {
	f := make([]int, N)
	gdns := make([]map[int]int, N)
	for i := 0; i < N; i++ {
		gdns[i] = make(map[int]int)
	}
	for i := range paths {
		gdns[paths[i][0]-1][paths[i][1]] = 1
		gdns[paths[i][1]-1][paths[i][0]] = 1
	}
	var adjFlag bool
	for i := range gdns {
		for j := 1; j <= 4; j++ {
			adjFlag = false
			for g := range gdns[i] {
				if f[g-1] == j {
					adjFlag = true
					break
				}
			}
			if !adjFlag {
				f[i] = j
				break
			}
		}
	}
	return f
}
