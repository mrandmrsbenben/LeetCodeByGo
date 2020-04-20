package main

import "fmt"

func main() {
	// nums := []int{4, 5, 6, 8, 0, 1, 2}
	// nums := []int{4, 5, 6, 8, 1, 2, 3}
	// nums := []int{5, 1, 2, 3, 4}
	nums := []int{5, 1}
	fmt.Println("Output: ", search(nums, 5))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.6 MB, 在所有 Go 提交中击败了100.00%的用户
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	if l == r && nums[0] == target {
		return 0
	}

	var m int
	for l < r {
		m = l + (r-l)/2
		switch target {
		case nums[m]:
			return m
		case nums[l]:
			return l
		case nums[r]:
			return r
		default:
			if nums[l] < nums[r] {
				// not rotated
				if target < nums[l] || target > nums[r] {
					return -1
				} else if target > nums[l] && target < nums[m] {
					r = m
				} else {
					l = m + 1
				}
			} else {
				// rotated
				if target < nums[l] && target > nums[r] {
					return -1
				}

				if nums[l] < nums[m] {
					if target > nums[l] && target < nums[m] {
						r = m
					} else {
						l = m + 1
					}
				} else if target < nums[m] || target > nums[l] {
					r = m
				} else if target > nums[m] && target < nums[l] {
					l = m + 1
				} else if target > nums[m] && target > nums[l] {
					r = m
				}
			}
		}
	}
	return -1
}
