package main

import "fmt"

func main() {
	// coordinates := [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}
	coordinates := [][]int{{3, 4}, {1, 2}, {2, 3}, {4, 5}, {5, 6}, {6, 7}, {-1, 0}}
	fmt.Println("Output: ", checkStraightLine(coordinates))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了95.45%的用户
//内存消耗 :3.7 MB, 在所有 Go 提交中击败了30.00%的用户
func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true
	}

	if coordinates[0][0] == coordinates[1][0] {
		for i := 2; i < len(coordinates); i++ {
			if coordinates[i][0] != coordinates[0][0] {
				return false
			}
		}
	} else {
		k := float64(coordinates[1][1]-coordinates[0][1]) / float64(coordinates[1][0]-coordinates[0][0])
		for i := 2; i < len(coordinates); i++ {
			if coordinates[i][0] == coordinates[0][0] {
				return false
			} else if k != float64(coordinates[i][1]-coordinates[0][1])/float64(coordinates[i][0]-coordinates[0][0]) {
				return false
			}
		}
	}

	return true
}
