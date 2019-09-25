package main

import "fmt"

func main() {
	calories := []int{1, 2, 3, 4, 5}
	k := 2
	lower := 3
	upper := 6
	fmt.Println("Output: ", dietPlanPerformance(calories, k, lower, upper))
	// calories := []int{3, 2}
	// k := 2
	// lower := 0
	// upper := 1
	// fmt.Println("Output: ", dietPlanPerformance(calories, k, lower, upper))
	// calories := []int{6, 5, 0, 0}
	// k := 2
	// lower := 1
	// upper := 5
	// fmt.Println("Output: ", dietPlanPerformance(calories, k, lower, upper))
}

//执行用时 :28 ms, 在所有 Go 提交中击败了70.45%的用户
//内存消耗 :8.2 MB, 在所有 Go 提交中击败了100.00%的用户
func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	points := 0
	var T int
	for i := 0; i < k-1; i++ {
		T += calories[i]
	}
	for i := k - 1; i < len(calories); i++ {
		T += calories[i]
		if T < lower {
			points--
		} else if T > upper {
			points++
		}
		T -= calories[i-k+1]
	}
	return points
}

//执行用时 :1580 ms, 在所有 Go 提交中击败了11.36%的用户
//内存消耗 :8.2 MB, 在所有 Go 提交中击败了100.00%的用户
func dietPlanPerformanceV1(calories []int, k int, lower int, upper int) int {
	points := 0
	var T int
	for i := 0; i < len(calories)-k+1; i++ {
		T = 0
		for j := 0; j < k; j++ {
			T += calories[i+j]
		}
		if T < lower {
			points--
		} else if T > upper {
			points++
		}
	}
	return points
}
