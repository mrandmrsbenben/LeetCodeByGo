package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/
func main() {
	// obj := Constructor()
	// obj.AddNum(1)
	// obj.AddNum(2)
	// fmt.Println("Output 1: ", obj.FindMedian())
	// obj.AddNum(3)
	// fmt.Println("Output 2: ", obj.FindMedian())
	obj := Constructor()
	obj.AddNum(-1)
	fmt.Println("Output 1: ", obj.FindMedian())
	obj.AddNum(-2)
	fmt.Println("Output 2: ", obj.FindMedian())
	obj.AddNum(-3)
	fmt.Println("Output 3: ", obj.FindMedian())
	obj.AddNum(-4)
	fmt.Println("Output 4: ", obj.FindMedian())
	obj.AddNum(-5)
	fmt.Println("Output 5: ", obj.FindMedian())
}

//执行用时 :356 ms, 在所有 Go 提交中击败了18.60%的用户
//内存消耗 :155.7 MB, 在所有 Go 提交中击败了100.00%的用户
type MedianFinder struct {
	data []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{[]int{}}
}

func (this *MedianFinder) AddNum(num int) {
	switch len(this.data) {
	case 0:
		this.data = []int{num}
	case 1:
		if num < this.data[0] {
			this.data = []int{num, this.data[0]}
		} else {
			this.data = []int{this.data[0], num}
		}
	case 2:
		if num < this.data[0] {
			this.data = []int{num, this.data[0], this.data[1]}
		} else if num > this.data[1] {
			this.data = []int{this.data[0], this.data[1], num}
		} else {
			this.data = []int{this.data[0], num, this.data[1]}
		}
	default:
		this.data = insert(this.data, num)
	}
}

func insert(data []int, num int) []int {
	if num < data[0] {
		return append([]int{num}, data...)
	} else if num > data[len(data)-1] {
		return append(data, num)
	}
	index := sort.SearchInts(data, num)
	// index := -1
	// l, r := 0, len(data)-1
	// for {
	// 	index = l + (r-l)/2
	// 	if num == data[index] {
	// 		break
	// 	} else if num < data[index] {
	// 		r = index
	// 	} else {
	// 		l = index
	// 	}
	// 	if r-l <= 1 {
	// 		index = r
	// 		break
	// 	}
	// }
	return append(data[0:index], append([]int{num}, data[index:]...)...)
}

func (this *MedianFinder) FindMedian() float64 {
	var res float64
	l := len(this.data)
	if l%2 == 0 {
		res = float64(this.data[l/2]+this.data[l/2-1]) / 2
	} else {
		res = float64(this.data[l/2])
	}
	return res
}
