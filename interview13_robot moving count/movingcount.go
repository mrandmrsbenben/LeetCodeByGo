package main

import "fmt"

func main() {
	m := 38
	n := 15
	k := 9
	fmt.Println("Output: ", movingCount(m, n, k))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.7 MB, 在所有 Go 提交中击败了100.00%的用户
func movingCount(m int, n int, k int) int {
	cnt := 0

	moves := make([][]int, m)
	for i := 0; i < m; i++ {
		moves[i] = make([]int, n)
	}

	var nextstep func(x, y int)
	nextstep = func(x, y int) {
		if sum(x)+sum(y) > k {
			return
		}
		cnt++
		moves[x][y] = 1
		if x > 0 && moves[x-1][y] == 0 {
			nextstep(x-1, y)
		}
		if x+1 < m && moves[x+1][y] == 0 {
			nextstep(x+1, y)
		}
		if y > 0 && moves[x][y-1] == 0 {
			nextstep(x, y-1)
		}
		if y+1 < n && moves[x][y+1] == 0 {
			nextstep(x, y+1)
		}
	}

	nextstep(0, 0)

	return cnt
}

func sum(n int) int {
	res := 0
	for n > 0 {
		res += n % 10
		n = n / 10
	}
	return res
}
