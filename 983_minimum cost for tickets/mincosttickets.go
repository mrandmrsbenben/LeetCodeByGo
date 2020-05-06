package main

import "fmt"

func main() {
	// days := []int{1, 4, 6, 7, 8, 20}
	// costs := []int{2, 7, 15}
	// days := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}
	// costs := []int{2, 7, 15}
	days := []int{1, 2, 4, 8, 12, 14, 15, 16, 22, 23, 25, 26, 28,
		29, 32, 36, 38, 40, 41, 44, 46, 49, 51, 54, 56, 58, 64, 65,
		66, 67, 68, 69, 72, 73, 74, 75, 76, 78, 80, 82, 83, 86, 87,
		88, 91, 94, 96, 98, 99}
	costs := []int{5, 27, 93}
	fmt.Println("Output: ", mincostTickets(days, costs))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.4 MB, 在所有 Go 提交中击败了100.00%的用户
func mincostTickets(days []int, costs []int) int {
	buf := make([]int, len(days))

	var findMin func([]int) int
	findMin = func(d []int) int {
		if len(d) == 0 {
			return 0
		}
		i := len(days) - len(d)
		if buf[i] > 0 {
			return buf[i]
		}
		cost := min(costs[0]+findMin(d[1:]),
			costs[1]+findMin(afterExpireDay(d, 7)),
			costs[2]+findMin(afterExpireDay(d, 30)))
		buf[i] = cost
		return cost
	}

	return findMin(days)
}

func afterExpireDay(days []int, length int) []int {
	res := []int{}
	if len(days) == 0 {
		return res
	}
	expireDay := days[0] + length
	for i := 1; i < len(days); i++ {
		if days[i] >= expireDay {
			res = days[i:]
			break
		}
	}
	return res
}

func min(x, y, z int) int {
	if x < y {
		if z < x {
			return z
		}
		return x
	}
	if y < z {
		return y
	}
	return z
}
