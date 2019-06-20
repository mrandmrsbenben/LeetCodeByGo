package main

import (
	"fmt"
	"sort"
)

func main() {
	costs := [][]int{{10, 20}, {60, 70}, {10, 20}, {30, 20}}
	fmt.Printf("Output: %d\n", twoCitySchedCost(costs))
}

func twoCitySchedCost(costs [][]int) int {
	cost := 0
	c := make([]int, 0)
	m := make(map[int][][]int)
	for _, v := range costs {
		if cs, ok := m[v[1]-v[0]]; ok {
			m[v[1]-v[0]] = append(cs, v)
		} else {
			m[v[1]-v[0]] = [][]int{v}
		}
		c = append(c, v[1]-v[0])
	}
	sort.Ints(c)
	count := 0
	for _, v := range c {
		if len(m[v]) == 1 {
			if count < len(costs)/2 {
				cost = cost + m[v][0][1]
			} else {
				cost = cost + m[v][0][0]
			}
			count = count + 1
		} else {
			if count < len(costs)/2 {
				cost = cost + m[v][0][1]
			} else {
				cost = cost + m[v][0][0]
			}
			count = count + 1
			m[v] = m[v][1:]
		}
	}
	return cost
}
