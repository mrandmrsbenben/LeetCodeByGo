package main

import "fmt"

func main() {
	rooms := [][]int{{1}, {2}, {3}, {}}
	// rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	fmt.Println("Output:", canVisitAllRooms(rooms))
}

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return false
	}
	rmap := make(map[int]int)
	rmap[0] = 1
	keys := []int{0}
	for len(keys) > 0 {
		l := len(keys)
		for _, i := range keys {
			for _, j := range rooms[i] {
				if rmap[j] != 1 {
					keys = append(keys, j)
					rmap[j] = 1
				}
			}
		}
		keys = keys[l:]
	}
	return len(rmap) == len(rooms)
}
