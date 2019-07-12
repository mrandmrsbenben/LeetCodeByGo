package main

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

//执行用时 :84 ms, 在所有 Go 提交中击败了93.62%的用户
//内存消耗 :10.6 MB, 在所有 Go 提交中击败了7.69%的用户
/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	nums := make([]int, len(this.nums))
	buf := make([]int, len(this.nums))
	copy(buf, this.nums)
	for i := range nums {
		j := rand.Intn(len(buf))
		nums[i] = buf[j]
		buf = append(buf[0:j], buf[j+1:]...)
	}
	return nums
}

//执行用时 :96 ms, 在所有 Go 提交中击败了82.98%的用户
//内存消耗 :9.4 MB, 在所有 Go 提交中击败了23.08%的用户
/** Returns a random shuffling of the array. */
func (this *Solution) ShuffleV1() []int {
	nums := make([]int, len(this.nums))
	copy(nums, this.nums)
	rand.Shuffle(len(nums), func(i int, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}
