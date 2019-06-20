package main

import "fmt"

//执行用时 :52 ms, 在所有 Go 提交中击败了49.30%的用户
//内存消耗 :7 MB, 在所有 Go 提交中击败了6.67%的用户
type MyLinkedList struct {
	nodes []int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{make([]int, 0)}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index >= 0 && index < len(this.nodes) {
		return this.nodes[index]
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.nodes = append([]int{val}, this.nodes...)
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.nodes = append(this.nodes, val)
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		this.AddAtHead(val)
	} else if index <= len(this.nodes) {
		if index == 0 {
			this.AddAtHead(val)
		} else if index == len(this.nodes) {
			this.AddAtTail(val)
		} else {
			buf := make([]int, len(this.nodes))
			copy(buf, this.nodes)
			this.nodes = append(append(buf[0:index], val), this.nodes[index:]...)
		}
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= 0 && index < len(this.nodes) {
		this.nodes = append(this.nodes[0:index], this.nodes[index+1:]...)
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
func main() {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	fmt.Println(linkedList.nodes)
	linkedList.AddAtTail(3)
	fmt.Println(linkedList.nodes)
	linkedList.AddAtIndex(1, 2)
	fmt.Println(linkedList.nodes)
	fmt.Println(linkedList.Get(1))
	linkedList.DeleteAtIndex(1)
	fmt.Println(linkedList.nodes)
	fmt.Println(linkedList.Get(1))
}
