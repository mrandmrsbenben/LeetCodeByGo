package main

import "fmt"

func main() {
	// intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	// newInterval := []int{4, 8}
	// intervals := [][]int{{1, 3}, {6, 9}, {12, 16}}
	// newInterval := []int{2, 5}
	// intervals := [][]int{{1, 5}, {6, 9}, {12, 16}}
	// // newInterval := []int{2, 3}
	// intervals := [][]int{{1, 5}}
	// newInterval := []int{0, 0}
	// intervals := [][]int{{0, 5}, {8, 9}}
	// newInterval := []int{3, 4}
	intervals := [][]int{{3, 5}, {8, 9}}
	newInterval := []int{6, 6}
	fmt.Println("Output: ", insert(intervals, newInterval))
}

//执行用时：8 ms, 在所有 Go 提交中击败了83.73%的用户
//内存消耗：4.8 MB, 在所有 Go 提交中击败了38.58%的用户
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	res := [][]int{}
	inserted := false
	for _, t := range intervals {
		if t[1] < newInterval[0] || t[0] > newInterval[1] {
			res = append(res, t)
			continue
		}
		if t[0] < newInterval[0] {
			newInterval[0] = t[0]
		}
		if t[0] > newInterval[1] {
			res = append(res, newInterval)
			inserted = true
		} else if t[1] >= newInterval[1] {
			newInterval[1] = t[1]
			res = append(res, newInterval)
			inserted = true
		}
	}

	if !inserted {
		if len(res) == 0 {
			return [][]int{newInterval}
		} else if newInterval[1] < res[0][0] {
			res = append([][]int{newInterval}, res...)
		} else if newInterval[0] > res[len(res)-1][1] {
			res = append(res, newInterval)
		} else {
			for i, t := range res {
				if t[1] > newInterval[0] {
					res = append(res[:i], append([][]int{newInterval}, res[i:]...)...)
					break
				}
			}
		}
	}
	return res
}
