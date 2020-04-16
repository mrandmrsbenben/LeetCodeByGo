package main

import (
	"fmt"
	"sort"
)

func main() {
	// intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals := [][]int{{1, 4}, {1, 4}}
	// intervals := [][]int{{2, 3}, {5, 5}, {2, 2}, {3, 4}}
	// intervals := [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	fmt.Println("Output: ", merge(intervals))
}

//执行用时 :12 ms, 在所有 Go 提交中击败了79.14%的用户
//内存消耗 :4.7 MB, 在所有 Go 提交中击败了100.00%的用户
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if cur[0] == intervals[i][0] {
			cur[1] = intervals[i][1]
		} else if cur[1] >= intervals[i][0] && cur[1] <= intervals[i][1] {
			cur[1] = intervals[i][1]
		} else if cur[1] < intervals[i][0] {
			res = append(res, cur)
			cur = intervals[i]
		}
	}
	res = append(res, cur)
	return res
}

//执行用时 :32 ms, 在所有 Go 提交中击败了18.71%的用户
//内存消耗 :6.4 MB, 在所有 Go 提交中击败了11.11%的用户
func mergeV1(intervals [][]int) [][]int {
	buf := [][]int{}
	bm := make(map[int][]int)
	flags := make([]bool, len(intervals))
	merged := true
	for len(intervals) > 0 {
		merged = false
		for i := 0; i < len(intervals); i++ {
			if flags[i] {
				continue
			}
			flags[i] = true
			for j := i + 1; j < len(intervals); j++ {
				if flags[j] {
					continue
				}
				if intervals[i][1] < intervals[j][0] || intervals[i][0] > intervals[j][1] {
					// not merge
					if _, ok := bm[j]; !ok {
						buf = append(buf, intervals[j])
						bm[j] = intervals[j]
					}
				} else if intervals[i][0] <= intervals[j][0] && intervals[i][1] >= intervals[j][1] {
					// merge as i (delete j)
					flags[j] = true
					merged = true
				} else if intervals[i][0] >= intervals[j][0] && intervals[i][1] <= intervals[j][1] {
					// merge as j, break loop
					delete(bm, i)
					intervals[i] = intervals[j]
					flags[j] = true
					merged = true
				} else if intervals[i][0] <= intervals[j][0] && intervals[i][1] <= intervals[j][1] {
					// merge as {i[0],j[1]}, break loop
					intervals[i][1] = intervals[j][1]
					flags[j] = true
					merged = true
				} else if intervals[i][0] >= intervals[j][0] && intervals[i][1] >= intervals[j][1] {
					// merge as {j[0],i[1]}, break loop
					intervals[i][0] = intervals[j][0]
					flags[j] = true
					merged = true
				}
			}
			if _, ok := bm[i]; !ok {
				buf = append(buf, intervals[i])
				bm[i] = intervals[i]
			}
		}
		if !merged {
			break
		}
		intervals = make([][]int, len(buf))
		copy(intervals, buf)
		flags = make([]bool, len(intervals))
		buf = [][]int{}
		bm = make(map[int][]int)
	}
	return intervals
}

func getIntervals(bm map[int][]int) [][]int {
	res := make([][]int, len(bm))
	i := 0
	for n := range bm {
		res[i] = bm[n]
		i++
	}
	return res
}
