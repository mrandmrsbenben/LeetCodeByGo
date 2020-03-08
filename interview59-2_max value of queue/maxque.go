package main

import "fmt"

// https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/
func main() {
	mq := Constructor()
	mq.Push_back(1)
	mq.Push_back(2)
	fmt.Println("Output 1: ", mq.Max_value())
	fmt.Println("Output 2: ", mq.Pop_front())
	fmt.Println("Output 3: ", mq.Max_value())
}

type MaxQueue struct {
	que  []int
	maxq []int
}

func Constructor() MaxQueue {
	return MaxQueue{[]int{}, []int{}}
}

func (this *MaxQueue) Max_value() int {
	if len(this.maxq) == 0 {
		return -1
	}
	return this.maxq[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.que = append(this.que, value)
	for len(this.maxq) > 0 && this.maxq[len(this.maxq)-1] < value {
		this.maxq = this.maxq[0 : len(this.maxq)-1]
	}
	this.maxq = append(this.maxq, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.que) == 0 {
		return -1
	}
	if len(this.maxq) > 0 && this.maxq[0] == this.que[0] {
		this.maxq = this.maxq[1:]
	}
	value := this.que[0]
	this.que = this.que[1:]
	return value
}
