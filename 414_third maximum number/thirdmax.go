package main

import "fmt"

func main() {
	// common.MakeTestCases("thirdMax")
	// input1 := []int{3, 2, 1}
	// fmt.Printf("Expect1: 1\n")
	// fmt.Printf("Output1: %v\n", thirdMax(input1))
	// input2 := []int{1, 2}
	// fmt.Printf("Expect2: 2\n")
	// fmt.Printf("Output2: %v\n", thirdMax(input2))
	// input3 := []int{2, 2, 3, 1}
	// fmt.Printf("Expect3: 1\n")
	// fmt.Printf("Output3: %v\n", thirdMax(input3))
	input4 := []int{1, 2, 2}
	fmt.Printf("Output4: %v\n", thirdMax(input4))
}

//执行用时 :8 ms, 在所有Go提交中击败了96.88%的用户
//内存消耗 :3 MB, 在所有Go提交中击败了83.33%的用户
func thirdMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	max1, max2 := 0, 0
	max3 := 0
	mflag := []bool{false, false, false}
	for i := 0; i < len(nums); i++ {
		if mflag[0] == false {
			max1 = nums[i]
			mflag[0] = true
		} else if mflag[1] == false && nums[i] != max1 {
			if nums[i] > max1 {
				max1, max2 = nums[i], max1
			} else {
				max2 = nums[i]
			}
			mflag[1] = true
		} else if mflag[1] && mflag[2] == false && nums[i] != max2 && nums[i] != max1 {
			if nums[i] > max1 {
				max3, max2 = max2, max1
				max1 = nums[i]
			} else if nums[i] > max2 {
				max2, max3 = nums[i], max2
			} else {
				max3 = nums[i]
			}
			mflag[2] = true
		} else if mflag[0] && mflag[1] && mflag[2] {
			if nums[i] > max1 {
				max3, max2 = max2, max1
				max1 = nums[i]
			} else if nums[i] > max2 && nums[i] < max1 {
				max2, max3 = nums[i], max2
			} else if nums[i] > max3 && nums[i] < max2 {
				max3 = nums[i]
			}
		}
	}
	if mflag[2] {
		return max3
	}
	return max1
}
