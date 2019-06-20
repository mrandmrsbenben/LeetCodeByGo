package main

import (
	"fmt"
	"sort"
)

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())

	// s := []int{0, 1, 2}
	// fmt.Println(s[0:0], s[3:])
}

type MinStack struct {
	s      []int
	sorted []int
}

func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.s = append(this.s, x)
	index := sort.SearchInts(this.sorted, x)
	switch index {
	case 0:
		this.sorted = append([]int{x}, this.sorted...)
	case len(this.sorted):
		this.sorted = append(this.sorted, x)
	default:
		cp := make([]int, len(this.sorted))
		copy(cp, this.sorted)
		this.sorted = append(append(cp[0:index], x), this.sorted[index:]...)
	}
}

func (this *MinStack) Pop() {
	if len(this.s) > 0 {
		index := sort.SearchInts(this.sorted, this.s[len(this.s)-1])
		switch index {
		case 0:
			this.sorted = this.sorted[1:]
		case len(this.sorted) - 1:
			this.sorted = this.sorted[0 : len(this.sorted)-1]
		default:
			this.sorted = append(this.sorted[0:index], this.sorted[index+1:]...)
		}
		this.s = this.s[0 : len(this.s)-1]
	}
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	return this.sorted[0]
}
