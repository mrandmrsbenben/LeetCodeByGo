package main

import (
	"fmt"
)

// Node Node
type Node struct {
	key, val   int
	freq       int
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

// LFUCache LFU cache
type LFUCache struct {
	vals   map[int]*Node
	cap    int
	lists  map[int]*DList
	minfrq int
}

// Constructor LFU Cache Constructor
func Constructor(capacity int) LFUCache {
	vals := make(map[int]*Node)
	lists := make(map[int]*DList)
	lists[1] = createList()
	return LFUCache{vals, capacity, lists, 1}
}

// Get get value by key from LFUCache
func (this *LFUCache) Get(key int) int {
	node := this.vals[key]
	if node == nil {
		return -1
	}
	this.increaseFreq(node)
	return node.val
}

// Put put value into LFUCache
func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}

	if node, ok := this.vals[key]; ok {
		node.val = value
		this.increaseFreq(node)
	} else if this.cap > len(this.vals) {
		node := &Node{key, value, 1, nil, nil}
		this.vals[key] = node
		list := this.lists[1]
		list.add(node)
		if this.minfrq > 1 {
			this.minfrq = 1
		}
	} else {
		list := this.lists[this.minfrq]
		node = list.pop()
		delete(this.vals, node.key)

		node = &Node{key, value, 1, nil, nil}
		this.vals[key] = node
		this.lists[1].add(node)
		this.minfrq = 1
	}
}

func (this *LFUCache) increaseFreq(node *Node) {
	oldlist := this.lists[node.freq]
	oldlist.remove(node)
	if this.minfrq == node.freq && oldlist.isEmpty {
		this.minfrq++
	}
	node.freq++
	if newlist, ok := this.lists[node.freq]; ok {
		newlist.add(node)
	} else {
		newlist = createList()
		newlist.add(node)
		this.lists[node.freq] = newlist
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// LFUCache cache = new LFUCache( 2 /* capacity */ )
// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // returns 1
// cache.put(3, 3);    // evicts key 2
// cache.get(2);       // returns -1 (not found)
// cache.get(3);       // returns 3.
// cache.put(4, 4);    // evicts key 1.
// cache.get(1);       // returns -1 (not found)
// cache.get(3);       // returns 3
// cache.get(4);       // returns 4

//执行用时 :144 ms, 在所有 Go 提交中击败了68.63%的用户
//内存消耗 :17 MB, 在所有 Go 提交中击败了100.00%的用户
func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println("Get 1:", cache.Get(1))
	cache.Put(3, 3)
	fmt.Println("Get 2:", cache.Get(2))
	fmt.Println("Get 3:", cache.Get(3))
	cache.Put(4, 4)
	fmt.Println("Get 1:", cache.Get(1))
	fmt.Println("Get 3:", cache.Get(3))
	fmt.Println("Get 4:", cache.Get(4))
}
