package main

import "fmt"

type MyHashMap struct {
	m []int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	m := make([]int, 1000000)
	for i := range m {
		m[i] = -1
	}
	return MyHashMap{m}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.m[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.m[key]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.m[key] = -1
}

func main() {
	hashMap := Constructor()
	hashMap.Put(1, 1)
	hashMap.Put(2, 2)
	fmt.Println(hashMap.Get(1)) // returns 1
	fmt.Println(hashMap.Get(3)) // returns -1 (not found)
	hashMap.Put(2, 1)           // update the existing value
	fmt.Println(hashMap.Get(2)) // returns 1
	hashMap.Remove(2)           // remove the mapping for 2
	fmt.Println(hashMap.Get(2)) // returns -1 (not found)
}
