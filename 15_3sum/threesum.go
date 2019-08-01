package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{0, 0, 0, 0}
	fmt.Println("Output:", threeSum(nums))
}

func threeSum(nums []int) [][]int {
	rs := make([][]int, 0)
	var l, r, sum int
	sort.Ints(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r = i+1, len(nums)-1
		for l < r {
			sum = nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				rs = append(rs, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return rs
}

func threeSumV1(nums []int) [][]int {
	rs := make([][]int, 0)
	nm := make(map[int]int, len(nums))
	sort.Ints(nums)
	for i, n := range nums {
		nm[n] = i
	}
	var c int
	for i := range nums {
		if nums[i] > 0 {
			break
		} else if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] > 0 {
				break
			} else if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			c = -1 * (nums[i] + nums[j])
			if k, ok := nm[c]; ok && k > j {
				rs = append(rs, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return rs
}

func threeSumSlow(nums []int) [][]int {
	rs := make([][]int, 0)
	sort.Ints(nums)
	for i := range nums {
		if nums[i] > 0 {
			break
		} else if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] > 0 {
				break
			} else if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				} else if nums[i]+nums[j]+nums[k] == 0 {
					rs = append(rs, []int{nums[i], nums[j], nums[k]})
					break
				}
			}
		}
	}
	return rs
}
