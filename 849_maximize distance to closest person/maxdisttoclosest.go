package main

import "fmt"

func main() {
	// seats := []int{1, 0, 0, 0, 1, 0, 1}
	seats := []int{1, 0, 0, 1}
	fmt.Println("Output:", maxDistToClosest(seats))
}

//执行用时 :12 ms, 在所有Go提交中击败了100.00%的用户
//内存消耗 :5.5 MB, 在所有Go提交中击败了66.67%的用户
func maxDistToClosest(seats []int) int {
	dist, max := 0, 0
	start := -1
	for i := range seats {
		if seats[i] == 1 {
			if start >= 0 {
				dist = (i - start) / 2
				if dist > max {
					max = dist
				}
			} else if i > 0 {
				dist = i
				if dist > max {
					max = dist
				}
			}
			start = i
		} else if i == len(seats)-1 {
			dist = i - start
			if dist > max {
				max = dist
			}
		}
	}
	return max
}
