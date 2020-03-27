package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("Output: ", nthUglyNumber(1690))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :4.1 MB, 在所有 Go 提交中击败了75.00%的用户
func nthUglyNumber(n int) int {
	res := make([]int, n)
	res[0] = 1

	var i2, i3, i5 int
	for i := 1; i < n; i++ {
		res[i] = min(res[i2]*2, res[i3]*3, res[i5]*5)
		if res[i] == res[i2]*2 {
			i2++
		}
		if res[i] == res[i3]*3 {
			i3++
		}
		if res[i] == res[i5]*5 {
			i5++
		}
	}
	return res[n-1]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

//执行用时 :64 ms, 在所有 Go 提交中击败了5.22%的用户
//内存消耗 :6.5 MB, 在所有 Go 提交中击败了100.00%的用户
func nthUglyNumberV1(n int) int {
	res := make([]int, n)
	m := make(map[int]int)

	h := &IntHeap{1}
	heap.Init(h)

	var k int
	for i := 0; i < n; i++ {
		k = heap.Pop(h).(int)
		res[i] = k
		if m[k*2] == 0 {
			heap.Push(h, k*2)
			m[k*2] = 1
		}
		if m[k*3] == 0 {
			heap.Push(h, k*3)
			m[k*3] = 1
		}
		if m[k*5] == 0 {
			heap.Push(h, k*5)
			m[k*5] = 1
		}
	}

	return res[n-1]
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
