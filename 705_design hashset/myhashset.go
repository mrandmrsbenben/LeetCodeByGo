package main

import "fmt"

type MyHashSet struct {
	set []bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{make([]bool, 1000000)}
}

func (this *MyHashSet) Add(key int) {
	this.set[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.set[key] = false
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.set[key]
}

func main() {
	hashSet := Constructor()
	hashSet.Add(1)
	hashSet.Add(2)
	fmt.Println(hashSet.Contains(1)) // returns true
	fmt.Println(hashSet.Contains(3)) // returns false (not found)
	hashSet.Add(2)
	fmt.Println(hashSet.Contains(2)) // returns true
	hashSet.Remove(2)
	fmt.Println(hashSet.Contains(2)) // returns false (already removed)
}
