package main

import (
	"fmt"
	"sort"
)

func main() {
	// people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	people := [][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}
	fmt.Println("Output:", reconstructQueue(people))
}

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
