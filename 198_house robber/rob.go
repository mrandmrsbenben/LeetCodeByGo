package main

import (
	"fmt"
	"sort"
)

func main() {
	// common.MakeTestCases("rob")
	// input1 := []int{1, 2, 3, 1}
	// fmt.Printf("Expect1: 4\n")
	// fmt.Printf("Output1: %v\n", rob(input1))
	// input2 := []int{2, 7, 9, 3, 1}
	// fmt.Printf("Expect2: 12\n")
	// fmt.Printf("Output2: %v\n", rob(input2))
	// input3 := []int{4, 9, 3, 4}
	// fmt.Printf("Output3: %v\n", rob(input3))
	input4 := []int{226, 174, 214, 16, 218, 48, 153, 131, 128, 17, 157, 142, 88, 43, 37, 157, 43,
		221, 191, 68, 206, 23, 225, 82, 54, 118, 111, 46, 80, 49, 245, 63, 25, 194, 72, 80, 143,
		209, 18, 55, 122, 65, 66, 177, 101, 63, 201, 172, 130, 103, 225, 142, 46, 86, 185, 62, 138,
		212, 192, 125, 77, 223, 188, 99, 228, 90, 25, 193, 211, 84, 239, 119, 234, 85, 83, 123, 120,
		131, 203, 219, 10, 82, 35, 120, 180, 249, 106, 37, 169, 225, 54, 103, 55, 166, 124}
	fmt.Printf("Output4: %v\n", rob(input4))
}

func robV2(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	rob0 := robSum(nums)
	rob1 := robSum(nums[1:])
	// rob2 := nums[0] + robSum(nums[2:])
	robs := []int{rob0, rob1}
	// rob6 := nums[0] + robSum(nums[2:len(nums)-2]) + nums[len(nums)-1]
	// rob7 := nums[1] + robSum(nums[3:len(nums)-2]) + nums[len(nums)-1]
	// reverse(nums)
	// rob3 := robSum(nums)
	// rob4 := robSum(nums[1:])
	// rob5 := nums[0] + robSum(nums[2:])
	// robs := []int{rob0, rob1, rob2, rob3, rob4, rob5}
	fmt.Println(robs)
	sort.Ints(robs)
	return robs[len(robs)-1]
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

func robSumE4(nums []int) int {
	rob := 0
	var sum, steps int
	for {
		if len(nums) < 4 {
			sum, steps = robEvery4(nums[0:])
		} else {
			sum, steps = robEvery4(nums[0:4])
		}
		rob += sum
		if steps == 0 || steps+1 > len(nums) {
			break
		}
		nums = nums[steps+1:]
	}
	return rob
}

func robSumV2(nums []int) int {
	rob := 0
	var sum, steps int
	for {
		if len(nums) < 3 {
			sum, steps = robEvery3(nums[0:])
		} else {
			sum, steps = robEvery3(nums[0:3])
		}
		rob += sum
		if steps == 0 || steps+1 > len(nums) {
			break
		}
		nums = nums[steps+1:]
	}
	return rob
}

func robEvery3(nums []int) (int, int) {
	switch len(nums) {
	case 0:
		return 0, 0
	case 1:
		return nums[0], 0
	case 2:
		if nums[0] > nums[1] {
			return nums[0], 0
		}
		return nums[1], 0
	default:
		if nums[0]+nums[2] > nums[1] {
			return nums[0] + nums[2], 3
		}
		return nums[1], 2
	}
}

func robEvery4(nums []int) (int, int) {
	switch len(nums) {
	case 0:
		return 0, 0
	case 1:
		return nums[0], 0
	case 2:
		if nums[0] > nums[1] {
			return nums[0], 0
		}
		return nums[1], 0
	case 3:
		if nums[0]+nums[2] > nums[1] {
			return nums[0] + nums[2], 0
		}
		return nums[1], 0
	default:
		sum0 := nums[0] + nums[2]
		sum1 := nums[0] + nums[3]
		sum2 := nums[1] + nums[3]
		if sum0 > sum1 {
			if sum2 > sum0 {
				return sum2, 4
			}
			return sum0, 3
		}
		if sum2 > sum1 {
			return sum2, 4
		}
		return sum1, 4
	}
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	var max0, max1 int
	if len(nums) > 10 {
		max0 = longSum(nums)
		max1 = longSum(nums[1:])
	} else {
		max0 = robSum(nums)
		max1 = robSum(nums[1:])
	}
	if max1 > max0 {
		return max1
	}
	return max0
}

func longSum(nums []int) int {
	sum0 := robSum(nums[0:len(nums)/2]) + robSum(nums[len(nums)/2+1:])
	sum1 := robSum(nums[0:len(nums)/2-1]) + robSum(nums[len(nums)/2:])
	if sum0 > sum1 {
		return sum0
	}
	return sum1
}

func robSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	sum0 := nums[0] + robSum(nums[2:])
	sum1 := nums[1] + robSum(nums[3:])
	if sum0 > sum1 {
		return sum0
	}
	return sum1
}
