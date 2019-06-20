package main

import "fmt"

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i > len(this.nums)-1 {
		return 0
	}
	if i < 0 {
		i = 0
	}
	if j > len(this.nums)-1 {
		j = len(this.nums) - 1
	}
	sum := 0
	// nums := this.nums[i : j+1]
	// for _, v := range nums {
	// 	sum = sum + v
	// }
	for n := i; n <= j; n++ {
		sum = sum + this.nums[n]
	}
	return sum
}

func main() {
	nums := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(nums.SumRange(0, 6))
	fmt.Println(nums.SumRange(2, 5))
	fmt.Println(nums.SumRange(0, 5))
}
