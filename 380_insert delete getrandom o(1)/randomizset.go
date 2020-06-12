package main

import (
	"fmt"
	"math/rand"
	"time"
)

//执行用时 :36 ms, 在所有 Go 提交中击败了58.23%的用户
//内存消耗 :7.7 MB, 在所有 Go 提交中击败了100.00%的用户
func main() {
	// ["RandomizedSet","insert","insert","remove","insert","remove","getRandom"]
	// [[],[0],[1],[0],[2],[1],[]]
	rs := Constructor()
	fmt.Println("Output: ", rs.Insert(0))
	fmt.Println("Output: ", rs.Insert(1))
	fmt.Println("Output: ", rs.Remove(0))
	fmt.Println("Output: ", rs.Insert(2))
	fmt.Println("Output: ", rs.Remove(1))
	fmt.Println("Output: ", rs.GetRandom())
}

// RandomizedSet RandomizedSet
type RandomizedSet struct {
	vals []int
	m    map[int]int
}

// Constructor Initialize data structure
func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano())
	return RandomizedSet{make([]int, 0), make(map[int]int)}
}

// Insert Inserts a value to the set. Returns true if the set did not already contain the specified element.
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.vals = append(this.vals, val)
	this.m[val] = len(this.vals) - 1
	return true
}

// Remove Removes a value from the set. Returns true if the set contained the specified element.
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}
	idx := this.m[val]
	delete(this.m, val)
	if idx != len(this.vals)-1 {
		this.vals[idx] = this.vals[len(this.vals)-1]
		this.m[this.vals[idx]] = idx
	}
	this.vals = this.vals[:len(this.vals)-1]
	return true
}

// GetRandom Get a random element from the set.
func (this *RandomizedSet) GetRandom() int {
	return this.vals[rand.Intn(len(this.vals))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
