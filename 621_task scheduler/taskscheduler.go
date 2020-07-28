package main

import (
	"fmt"
	"sort"
)

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	fmt.Println("Output: ", leastInterval(tasks, n))
}

//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：6 MB, 在所有 Go 提交中击败了100.00%的用户
func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	// Array Of Tasks By Letter
	c := make([]int, 26)
	for _, t := range tasks {
		c[t-'A']++
	}
	sort.Ints(c)

	max := c[25] - 1
	res := max * n
	for i := 24; i >= 0; i-- {
		if c[i] == 0 {
			continue
		}
		res -= min(c[i], max)
	}
	if res > 0 {
		return res + len(tasks)
	}
	return len(tasks)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
