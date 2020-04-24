package main

import "fmt"

func main() {
	// ["LRUCache","put","put","get","put","get","put","get","get","get"]
	// [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	// cache := Constructor(2)
	// cache.Put(1, 1)
	// cache.Put(2, 2)
	// fmt.Println("Output: ", cache.Get(1)) // returns 1
	// cache.Put(3, 3)                       // evicts key 2
	// fmt.Println("Output: ", cache.Get(2)) // returns -1 (not found)
	// cache.Put(4, 4)                       // evicts key 1
	// fmt.Println("Output: ", cache.Get(1)) // returns -1 (not found)
	// fmt.Println("Output: ", cache.Get(3)) // returns 3
	// fmt.Println("Output: ", cache.Get(4)) // returns 4

	// 	["LRUCache","put","put","put","put","get","get"]
	// [[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)
	fmt.Println("Output: ", cache.Get(1)) // returns 1
	fmt.Println("Output: ", cache.Get(2)) // returns -1 (not found)
}

// Node Node
type Node struct {
	key, val   int
	prev, next *Node
}

// DList Double direction List
type DList struct {
	head, tail *Node
	isEmpty    bool
}

func createList() *DList {
	return &DList{nil, nil, true}
}

func (this *DList) remove(node *Node) {
	if this.head == node && this.tail == node {
		this.head, this.tail = nil, nil
		this.isEmpty = true
	} else if this.head == node {
		this.head.next.prev = nil
		this.head = this.head.next
	} else if this.tail == node {
		this.tail.prev.next = nil
		this.tail = this.tail.prev
	} else {
		node.prev.next, node.next.prev = node.next, node.prev
		node.prev, node.next = nil, nil
	}

}

func (this *DList) add(node *Node) {
	if this.head == nil {
		this.head, this.tail = node, node
	} else {
		this.tail.next, node.prev = node, this.tail
		this.tail = node
	}
	this.isEmpty = false
}

func (this *DList) pop() *Node {
	node := this.head
	if this.head == nil {
		return nil
	} else if this.head.next == nil {
		this.head, this.tail = nil, nil
		this.isEmpty = true
		return node
	} else {
		this.head.next.prev = nil
		this.head = this.head.next
	}
	return node
}

//执行用时 :132 ms, 在所有 Go 提交中击败了89.82%的用户
//内存消耗 :16.7 MB, 在所有 Go 提交中击败了63.64%的用户
// LRUCache LRU cache
type LRUCache struct {
	vals map[int]*Node
	cap  int
	list *DList
}

// Constructor LRU Cache Constructor
func Constructor(capacity int) LRUCache {
	return LRUCache{make(map[int]*Node), capacity, createList()}
}

// Get get value by key from LRUCache
func (this *LRUCache) Get(key int) int {
	node := this.vals[key]
	if node == nil {
		return -1
	}
	this.list.remove(node)
	this.list.add(node)
	return node.val
}

// Put put value into LRUCache
func (this *LRUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}

	if node, ok := this.vals[key]; ok {
		node.val = value
		this.list.remove(node)
		this.list.add(node)
	} else if this.cap > len(this.vals) {
		node := &Node{key, value, nil, nil}
		this.vals[key] = node
		this.list.add(node)
	} else {
		node = this.list.pop()
		delete(this.vals, node.key)

		node = &Node{key, value, nil, nil}
		this.vals[key] = node
		this.list.add(node)
	}
}
