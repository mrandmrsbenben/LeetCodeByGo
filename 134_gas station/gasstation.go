package main

import "fmt"

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	// gas := []int{2, 3, 4}
	// cost := []int{3, 4, 3}
	fmt.Println("Output: ", canCompleteCircuit(gas, cost))
}

//执行用时：8 ms, 在所有 Go 提交中击败了36.32%的用户
//内存消耗：3 MB, 在所有 Go 提交中击败了17.69%的用户
func canCompleteCircuit(gas []int, cost []int) int {
	sum := 0
	tank := make([]int, len(gas))
	for i := range gas {
		tank[i] = gas[i] - cost[i]
		sum += tank[i]
	}
	if sum < 0 {
		return -1
	}

	for i := range tank {
		if tank[i] >= 0 {
			sum = tank[i]
			for j := i + 1; j < len(tank); j++ {
				sum += tank[j]
				if sum < 0 {
					break
				}
			}
			if sum < 0 {
				continue
			}
			for j := 0; j < i; j++ {
				sum += tank[j]
				if sum < 0 {
					break
				}
			}
			if sum >= 0 {
				return i
			}
		}
	}
	return -1
}

//执行用时：92 ms, 在所有 Go 提交中击败了23.32%的用户
//内存消耗：6.9 MB, 在所有 Go 提交中击败了5.38%的用户
func canCompleteCircuitV1(gas []int, cost []int) int {
	var g, c []int
	for i := range gas {
		if gas[i] >= cost[i] {
			g = append(gas[i:], gas[:i]...)
			c = append(cost[i:], cost[:i]...)
			if canComplete(g, c) {
				return i
			}
		}
	}
	return -1
}

func canComplete(gas []int, cost []int) bool {
	tank := 0
	for i := range gas {
		tank += gas[i] - cost[i]
		if tank < 0 {
			return false
		}
	}
	return true
}
