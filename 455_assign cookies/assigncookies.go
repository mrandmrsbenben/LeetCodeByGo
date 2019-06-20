package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{3, 2}
	s := []int{1}
	fmt.Printf("Output: %d\n", findContentChildren(g, s))
}

func findContentChildren(g []int, s []int) int {
	count := 0
	sort.Ints(g)
	sort.Ints(s)
	for _, child := range g {
		index := sort.SearchInts(s, child)
		if index == len(s) {
			break
		} else if index == 0 {
			s = s[1:]
		} else {
			s = append(s[0:index], s[index+1:]...)
		}
		count = count + 1
	}
	return count
}
