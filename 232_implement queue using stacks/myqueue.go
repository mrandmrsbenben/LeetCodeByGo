package main

import "fmt"

func main() {
	queue := &MyQueue{}
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek()) // returns 1
	fmt.Println(queue.Pop())  // returns 1
	fmt.Println(queue.Pop())  // returns 2
	fmt.Println(queue.Empty())
}

type MyQueue struct {
	l []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{make([]int, 0)}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.l = append(this.l, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	n := this.l[0]
	this.l = this.l[1:]
	return n
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.l[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.l) == 0
}
