package main

import "fmt"

//执行用时 :12 ms, 在所有 Go 提交中击败了98.68%的用户
//内存消耗 :6.3 MB, 在所有 Go 提交中击败了48.57%的用户
//MyCircularQueue MyCircularQueue
type MyCircularQueue struct {
	size  int
	queue []int
}

//Constructor Initialize your data structure here. Set the size of the queue to be k.
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{size: k, queue: make([]int, 0)}
}

//EnQueue Insert an element into the circular queue. Return true if the operation is successful.
func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.queue) == this.size {
		return false
	}
	this.queue = append(this.queue, value)
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if len(this.queue) == 0 {
		return false
	}
	this.queue = this.queue[1:]
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[0]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[len(this.queue)-1]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.queue) == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return len(this.queue) == this.size
}

func main() {
	circularQueue := Constructor(3)
	fmt.Println(circularQueue.IsEmpty())  // return true
	fmt.Println(circularQueue.EnQueue(1)) // return true
	fmt.Println(circularQueue.EnQueue(2)) // return true
	fmt.Println(circularQueue.EnQueue(3)) // return true
	fmt.Println(circularQueue.EnQueue(4)) // return false
	fmt.Println(circularQueue.Rear())     // return 3
	fmt.Println(circularQueue.IsFull())   // return true
	fmt.Println(circularQueue.DeQueue())  // return true
	fmt.Println(circularQueue.EnQueue(4)) // return true
	fmt.Println(circularQueue.Rear())     // return 4

}
