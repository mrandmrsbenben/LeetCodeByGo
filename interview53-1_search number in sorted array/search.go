package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	// nums := []int{-1}
	target := 8
	fmt.Println("Output: ", search(nums, target))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :4.1 MB, 在所有 Go 提交中击败了100.00%的用户
func search(nums []int, target int) int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		if nums[0] == target {
			return 1
		}
		return 0
	}

	fb, fe := false, false
	if target < nums[0] {
		fb = true
	}
	if target > nums[len(nums)-1] {
		fe = true
	}

	var i, j int
	l, r := []int{0, len(nums) - 1}, []int{0, len(nums) - 1}
	for !fb || !fe {
		if !fb {
			i = l[0] + (l[1]-l[0])/2
			if nums[i] < target {
				l[0] = i + 1
			} else if nums[i] == target && (i == 0 || (i > 0 && nums[i-1] < target)) {
				res[0] = i
				r[0] = i
				fb = true
			} else {
				l[1] = i
			}
			if l[0] == l[1] && nums[l[0]] != target {
				fb = true
			}
		}
		if !fe {
			j = r[0] + (r[1]-r[0])/2
			if nums[j] > target {
				r[1] = j
			} else if nums[j] == target && (j == len(nums)-1 || (j+1 < len(nums) && nums[j+1] > target)) {
				res[1] = j
				l[1] = j
				fe = true
			} else {
				r[0] = j + 1
			}
			if r[0] == r[1] && nums[r[0]] != target {
				fe = true
			}
		}
	}

	if res[0] == -1 {
		return 0
	}
	return res[1] - res[0] + 1
}
