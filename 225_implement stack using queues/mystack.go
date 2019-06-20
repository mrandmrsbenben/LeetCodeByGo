package main

import "fmt"

type MyStack struct {
	stack []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{make([]int, 0)}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	x := this.stack[len(this.stack)-1]
	this.stack = this.stack[0 : len(this.stack)-1]
	return x
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.stack[len(this.stack)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.stack) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
