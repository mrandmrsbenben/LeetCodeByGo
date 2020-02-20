package main

import "fmt"

//执行用时 :256 ms, 在所有 Go 提交中击败了27.27%的用户
//内存消耗 :8.4 MB, 在所有 Go 提交中击败了100.00%的用户
type Stack struct {
	data []int
}

func StackConstructor() Stack {
	return Stack{make([]int, 0)}
}

func (t *Stack) Push(n int) {
	t.data = append(t.data, n)
}

func (t *Stack) Pop() int {
	if len(t.data) == 0 {
		return -1
	}

	top := t.data[len(t.data)-1]
	t.data = t.data[0 : len(t.data)-1]
	return top
}

type CQueue struct {
	mainStack, subStack Stack
}

func Constructor() CQueue {
	return CQueue{StackConstructor(), StackConstructor()}
}

func (this *CQueue) AppendTail(value int) {
	this.mainStack.Push(value)
}

func (this *CQueue) DeleteHead() int {
	value := this.mainStack.Pop()
	for value != -1 {
		this.subStack.Push(value)
		value = this.mainStack.Pop()
	}
	ret := this.subStack.Pop()
	for ret != -1 {
		value = this.subStack.Pop()
		if value == -1 {
			break
		}
		this.mainStack.Push(value)
	}
	return ret
}

func main() {
	// obj := Constructor()
	// obj.AppendTail(3)
	// fmt.Println("Output 0: ", obj.DeleteHead())
	// fmt.Println("Output 1: ", obj.DeleteHead())

	obj := Constructor()
	fmt.Println("Output 0: ", obj.DeleteHead())
	obj.AppendTail(5)
	obj.AppendTail(2)
	fmt.Println("Output 1: ", obj.DeleteHead())
	fmt.Println("Output 1: ", obj.DeleteHead())
}
