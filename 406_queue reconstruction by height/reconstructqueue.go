package main

import (
	"fmt"
	"sort"
)

func main() {
	// people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	// people := [][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}
	people := [][]int{{40, 11}, {81, 12}, {32, 60}, {36, 39}, {76, 19},
		{11, 37}, {67, 13}, {45, 39}, {99, 0}, {35, 20}, {15, 3}, {62, 13},
		{90, 2}, {86, 0}, {26, 13}, {68, 32}, {91, 4}, {23, 24}, {73, 14},
		{86, 13}, {62, 6}, {36, 13}, {67, 9}, {39, 57}, {15, 45}, {37, 26},
		{12, 88}, {30, 18}, {39, 60}, {77, 2}, {24, 38}, {72, 7}, {96, 1},
		{29, 47}, {92, 1}, {67, 28}, {54, 44}, {46, 35}, {3, 85}, {27, 9},
		{82, 14}, {29, 17}, {80, 11}, {84, 10}, {5, 59}, {82, 6}, {62, 25},
		{64, 29}, {88, 8}, {11, 20}, {83, 0}, {94, 4}, {43, 42}, {73, 9},
		{57, 32}, {76, 24}, {14, 67}, {86, 2}, {13, 47}, {93, 1}, {95, 2},
		{87, 8}, {8, 78}, {58, 16}, {26, 75}, {26, 15}, {24, 56}, {69, 9},
		{42, 22}, {70, 17}, {34, 48}, {26, 39}, {22, 28}, {21, 8}, {51, 44},
		{35, 4}, {25, 48}, {78, 18}, {29, 30}, {13, 63}, {68, 8}, {21, 38},
		{56, 20}, {84, 14}, {56, 27}, {60, 40}, {98, 0}, {63, 7}, {27, 46},
		{70, 13}, {29, 23}, {49, 6}, {5, 64}, {67, 11}, {2, 31}, {59, 8},
		{93, 0}, {50, 39}, {84, 6}, {19, 39}}
	fmt.Println("Output:", reconstructQueue(people))
}

// func reconstructQueue(people [][]int) [][]int {
// 	sort.Slice(people, func(i int, j int) bool {
// 		if people[i][1] == people[j][1] {
// 			return people[i][0] < people[j][0]
// 		}
// 		return people[i][1] < people[j][1]
// 	})
// 	sort.Slice(people, func(i int, j int) bool {
// 		if people[i][1] == people[j][1] {
// 			return people[i][0] < people[j][0]
// 		}
// 		c := 0
// 		for k := 0; k < j; k++ {
// 			if people[k][0] >= people[i][0] {
// 				c++
// 			}
// 		}
// 		return c >= people[i][1]
// 	})
// 	return people
// }

//执行用时 :64 ms, 在所有 Go 提交中击败了64.00%的用户
//内存消耗 :8 MB, 在所有 Go 提交中击败了100.00%的用户
func reconstructQueue(people [][]int) [][]int {
	if len(people) <= 1 {
		return people
	}
	// group by k
	km := make(map[int][]int)
	for _, p := range people {
		if _, ok := km[p[1]]; ok {
			km[p[1]] = append(km[p[1]], p[0])
		} else {
			a := make([]int, 1)
			a[0] = p[0]
			km[p[1]] = a
		}
	}
	// sort every group
	ka := make([]int, len(km))
	var c int
	for k := range km {
		ka[c] = k
		sort.Ints(km[k])
		c++
	}
	sort.Ints(ka)
	// reconstruct array by group
	var h int
	rcq := make([][]int, 0)
	for _, k := range ka {
		for j := len(km[k]) - 1; j >= 0; j-- {
			h = km[k][j]
			c = 0
			for j := range rcq {
				if k == 0 {
					if h < rcq[j][0] {
						rcq = append(rcq[0:j], append([][]int{{h, k}}, rcq[j:]...)...)
						c++
						break
					}
				} else if rcq[j][0] >= h {
					c++
				}
				if c == k {
					rcq = append(rcq[0:j+1], append([][]int{{h, k}}, rcq[j+1:]...)...)
					break
				}
			}
			if c < k || len(rcq) == 0 {
				rcq = append(rcq, []int{h, k})
			}
		}
	}
	return rcq
}
