package main

import (
	"fmt"
)

//执行用时 :24 ms, 在所有 Go 提交中击败了56.95%的用户
//内存消耗 :8.2 MB, 在所有 Go 提交中击败了100.00%的用户
type MinStack struct {
	vals []int
	mins []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(x int) {
	this.vals = append(this.vals, x)
	if len(this.mins) == 0 || x <= this.mins[len(this.mins)-1] {
		this.mins = append(this.mins, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.vals) == 0 {
		return
	}
	x := this.vals[len(this.vals)-1]
	this.vals = this.vals[0 : len(this.vals)-1]
	if x == this.mins[len(this.mins)-1] {
		this.mins = this.mins[0 : len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	return this.vals[len(this.vals)-1]
}

func (this *MinStack) Min() int {
	return this.mins[len(this.mins)-1]
}

// https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/
func main() {
	minStack := Constructor()
	minStack.Push(0)
	minStack.Push(1)
	minStack.Push(0)
	fmt.Println(minStack.Min())
	minStack.Pop()
	// fmt.Println(minStack.Top())
	fmt.Println(minStack.Min())

}
