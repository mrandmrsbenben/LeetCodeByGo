package main

import "fmt"

func main() {
	// distance := []int{1, 2, 3, 4}
	// start := 0
	// destination := 3
	// fmt.Println("Output: ", distanceBetweenBusStops(distance, start, destination))
	distance := []int{3, 6, 7, 2, 9, 10, 7, 16, 11}
	start := 2
	destination := 6
	fmt.Println("Output: ", distanceBetweenBusStops(distance, start, destination))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.8 MB, 在所有 Go 提交中击败了100.00%的用户
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start == destination || len(distance) == 0 {
		return 0
	}

	dc, dcc := 0, 0
	if destination < start {
		for i := range distance {
			if i < destination || i >= start {
				dc += distance[i]
			} else {
				dcc += distance[i]
			}
		}
	} else {
		for i := range distance {
			if i < destination && i >= start {
				dc += distance[i]
			} else {
				dcc += distance[i]
			}
		}
	}

	if dc < dcc {
		return dc
	}
	return dcc
}
