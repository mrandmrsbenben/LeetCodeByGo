package main

import "fmt"

type RecentCounter struct {
	q []int
}

func Constructor() RecentCounter {
	q := make([]int, 0)
	return RecentCounter{q}
}

func (this *RecentCounter) Ping(t int) int {
	this.q = append(this.q, t)
	count := 0
	for i := len(this.q) - 1; i >= 0; i-- {
		if this.q[i] >= t-3000 && this.q[i] <= t {
			count = count + 1
		} else {
			break
		}
	}
	return count
}

func main() {
	rc := Constructor()
	pings := []int{1178, 1483, 1563, 4054, 4152}
	for _, t := range pings {
		fmt.Printf("Time: %d, Counter:%d\n", t, rc.Ping(t))
	}
}
