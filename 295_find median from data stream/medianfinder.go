package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// obj := Constructor()
	// obj.AddNum(1)
	// obj.AddNum(2)
	// fmt.Println("Output 1: ", obj.FindMedian())
	// obj.AddNum(3)
	// fmt.Println("Output 2: ", obj.FindMedian())
	obj := Constructor()
	obj.AddNum(-1)
	fmt.Println("Output 1: ", obj.FindMedian())
	obj.AddNum(-2)
	fmt.Println("Output 2: ", obj.FindMedian())
	obj.AddNum(-3)
	fmt.Println("Output 3: ", obj.FindMedian())
	obj.AddNum(-4)
	fmt.Println("Output 4: ", obj.FindMedian())
	obj.AddNum(-5)
	fmt.Println("Output 5: ", obj.FindMedian())
}

//执行用时 :172 ms, 在所有 Go 提交中击败了31.43%的用户
//内存消耗 :17.8 MB, 在所有 Go 提交中击败了56.25%的用户
type AscHeap []int

func (h AscHeap) Len() int           { return len(h) }
func (h AscHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h AscHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *AscHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *AscHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type DescHeap []int

func (h DescHeap) Len() int           { return len(h) }
func (h DescHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h DescHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *DescHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *DescHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	asc  AscHeap
	desc DescHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.asc, num)
	heap.Push(&this.desc, heap.Pop(&this.asc))
	if this.desc.Len() > this.asc.Len()+1 {
		heap.Push(&this.asc, heap.Pop(&this.desc))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.desc.Len() > this.asc.Len() {
		return float64(this.desc[0])
	}
	return (float64(this.desc[0]) + float64(this.asc[0])) / 2
}
