package main

import "fmt"

// https://leetcode-cn.com/problems/the-masseuse-lcci/
func main() {
	// nums := []int{1, 2, 3, 1}
	// nums := []int{2, 7, 9, 3, 1}
	// nums := []int{2, 1, 4, 5, 3, 1, 1, 3}
	nums := []int{226, 174, 214, 16, 218, 48, 153, 131, 128, 17, 157, 142, 88,
		43, 37, 157, 43, 221, 191, 68, 206, 23, 225, 82, 54, 118, 111, 46, 80,
		49, 245, 63, 25, 194, 72, 80, 143, 55, 209, 18, 55, 122, 65, 66, 177,
		101, 63, 201, 172, 130, 103, 225, 142, 46, 86, 185, 62, 138, 212, 192,
		125, 77, 223, 188, 99, 228, 90, 25, 193, 211, 84, 239, 119, 234, 85, 83,
		123, 120, 131, 203, 219, 10, 82, 35, 120, 180, 249, 106, 37, 169, 225,
		54, 103, 55, 166, 124}
	fmt.Println("Output: ", massage(nums))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.2 MB, 在所有 Go 提交中击败了100.00%的用户
func massage(nums []int) int {
	l := len(nums)
	m := make(map[int]int)

	var mass func(i int) int
	mass = func(i int) int {
		if n, ok := m[i]; ok {
			return n
		}

		max := 0
		switch l - i {
		case 0:
			return 0
		case 1:
			max = nums[i]
		case 2:
			if nums[i] > nums[i+1] {
				max = nums[i]
			} else {
				max = nums[i+1]
			}
		case 3:
			if nums[i]+nums[i+2] > nums[i+1] {
				max = nums[i] + nums[i+2]
			} else {
				max = nums[i+1]
			}
		default:
			m0 := nums[i] + mass(i+2)
			m1 := nums[i+1] + mass(i+3)
			if m0 > m1 {
				max = m0
			} else {
				max = m1
			}
		}
		m[i] = max
		return max
	}
	return mass(0)
}
