package main

import (
	"fmt"
)

func main() {
	// firstUnique := Constructor([]int{2, 3, 5})
	// fmt.Println("Output: ", firstUnique.ShowFirstUnique()) // return 2
	// firstUnique.Add(5)                                     // the queue is now [2,3,5,5]
	// fmt.Println("Output: ", firstUnique.ShowFirstUnique()) // return 2
	// firstUnique.Add(2)                                     // the queue is now [2,3,5,5,2]
	// fmt.Println("Output: ", firstUnique.ShowFirstUnique()) // return 3
	// firstUnique.Add(3)                                     // the queue is now [2,3,5,5,2,3]
	// fmt.Println("Output: ", firstUnique.ShowFirstUnique()) // return -1
	// firstUnique := Constructor([]int{7, 7, 7, 7, 7, 7})
	// fmt.Println("Output: ", firstUnique.ShowFirstUnique()) // return -1
	// firstUnique.Add(7)                                     // the queue is now [7,7,7,7,7,7,7]
	// firstUnique.Add(3)                                     // the queue is now [7,7,7,7,7,7,7,3]
	// firstUnique.Add(3)                                     // the queue is now [7,7,7,7,7,7,7,3,3]
	// firstUnique.Add(7)                                     // the queue is now [7,7,7,7,7,7,7,3,3,7]
	// firstUnique.Add(17)                                    // the queue is now [7,7,7,7,7,7,7,3,3,7,17]
	// fmt.Println("Output: ", firstUnique.ShowFirstUnique()) // return 17
	firstUnique := Constructor([]int{809})
	fmt.Println("Output: ", firstUnique.ShowFirstUnique()) // return 809
	firstUnique.Add(809)                                   // the queue is now [809,809]
	fmt.Println("Output: ", firstUnique.ShowFirstUnique()) // return -1

}

// TwoWayNode node of two way list
type TwoWayNode struct {
	val, count int
	prev, next *TwoWayNode
}

// TwoWayList two way list
type TwoWayList struct {
	nodes      map[int]*TwoWayNode
	head, tail *TwoWayNode
}

func createTwoWayList(nums []int) *TwoWayList {
	vals := make(map[int]*TwoWayNode)
	list := &TwoWayList{vals, nil, nil}
	for _, n := range nums {
		list.add(n)
	}
	return list
}

// add add node to list
func (t *TwoWayList) add(val int) {
	if node, ok := t.nodes[val]; !ok {
		// val is unique number
		node = &TwoWayNode{val, 1, nil, nil}
		t.nodes[val] = node
		if t.tail == nil {
			t.head, t.tail = node, node
		} else {
			node.prev, t.tail.next = t.tail, node
			t.tail = node
		}
	} else {
		// not a unique number
		if node.count == 1 {
			if t.head == t.tail && node == t.head {
				// only node
				t.head, t.tail = nil, nil
			} else if node == t.head {
				// remove head node
				t.head = t.head.next
			} else if node == t.tail {
				// remove tail node
				t.tail = t.tail.prev
			} else {
				// remove node from middle
				node.prev.next, node.next.prev = node.next, node.prev
			}
		}
		node.count++
	}
}

// pop return first node's value
func (t *TwoWayList) pop() int {
	if t.head == nil {
		return -1
	}
	return t.head.val
}

type FirstUnique struct {
	list *TwoWayList
}

func Constructor(nums []int) FirstUnique {
	return FirstUnique{createTwoWayList(nums)}
}

func (this *FirstUnique) ShowFirstUnique() int {
	return this.list.pop()
}

func (this *FirstUnique) Add(value int) {
	this.list.add(value)
}

/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */
