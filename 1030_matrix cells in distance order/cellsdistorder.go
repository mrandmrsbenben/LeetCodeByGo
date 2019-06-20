package main

import (
	"fmt"
	"strconv"
)

func main() {
	R := 1
	C := 100
	r0 := 0
	c0 := 35
	fmt.Printf("Output: %v\n", allCellsDistOrder(R, C, r0, c0))
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	cells := [][]int{{r0, c0}}
	m := make(map[string]int)
	var cur [][]int
	var key string
	prev := [][]int{{r0, c0}}
	m[strconv.Itoa(r0)+","+strconv.Itoa(c0)] = 1
	for {
		cur = make([][]int, 0)
		for _, v := range prev {
			if v[0]-1 >= 0 {
				key = strconv.Itoa(v[0]-1) + "," + strconv.Itoa(v[1])
				if _, ok := m[key]; !ok {
					cur = append(cur, []int{v[0] - 1, v[1]})
					m[key] = 1
				}
			}
			if v[1]-1 >= 0 {
				key = strconv.Itoa(v[0]) + "," + strconv.Itoa(v[1]-1)
				if _, ok := m[key]; !ok {
					cur = append(cur, []int{v[0], v[1] - 1})
					m[key] = 1
				}
			}
			if v[1]+1 < C {
				key = strconv.Itoa(v[0]) + "," + strconv.Itoa(v[1]+1)
				if _, ok := m[key]; !ok {
					cur = append(cur, []int{v[0], v[1] + 1})
					m[key] = 1
				}
			}
			if v[0]+1 < R {
				key = strconv.Itoa(v[0]+1) + "," + strconv.Itoa(v[1])
				if _, ok := m[key]; !ok {
					cur = append(cur, []int{v[0] + 1, v[1]})
					m[key] = 1
				}
			}
		}
		if len(cur) == 0 {
			break
		}
		cells = append(cells, cur...)
		prev = cur
	}
	return cells
}
