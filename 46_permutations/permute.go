package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println("Output:", permute(nums))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了99.62%的用户
//内存消耗 :6.9 MB, 在所有 Go 提交中击败了46.08%的用户
func permute(nums []int) [][]int {
	switch len(nums) {
	case 0:
		return [][]int{}
	case 1:
		return [][]int{nums}
	case 2:
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}
	l := 1
	for i := range nums {
		l *= (i + 1)
	}
	ps := make([][]int, l)
	var subs [][]int
	index := 0
	buf := make([]int, len(nums))
	for i := range nums {
		copy(buf, nums)
		if i == 0 {
			subs = permute(nums[i+1:])
		} else if i == len(nums)-1 {
			subs = permute(nums[0:i])
		} else {
			subs = permute(append(buf[0:i], buf[i+1:]...))
		}
		for j := range subs {
			ps[index] = append([]int{nums[i]}, subs[j]...)
			index++
		}
	}
	return ps
}

func permuteV0(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	} else if len(nums) == 1 {
		return [][]int{nums}
	}
	buf := make([]int, len(nums))
	copy(buf, nums)
	var buf2 []int
	pm := make(map[string][]int)
	for i := range nums {
		for j := range buf {
			if i != j {
				buf[i], buf[j] = buf[j], buf[i]
				buf2 = make([]int, len(buf))
				copy(buf2, buf)
				pm[toString(buf)] = buf2
			}
		}
		nums = append(nums[1:], nums[0])
		buf2 = make([]int, len(nums))
		copy(buf2, nums)
		pm[toString(nums)] = buf2
	}
	// for i := range nums {
	// 	copy(buf, nums)
	// 	for j := range buf {
	// 		if i != j {
	// 			buf[i], buf[j] = buf[j], buf[i]
	// 			buf2 = make([]int, len(buf))
	// 			copy(buf2, buf)
	// 			pm[toString(buf)] = buf2
	// 		}
	// 	}
	// }
	ps := make([][]int, len(pm))
	i := 0
	for _, p := range pm {
		ps[i] = p
		i++
	}
	return ps
}

func toString(nums []int) string {
	str := ""
	for _, n := range nums {
		str += strconv.Itoa(n) + ","
	}
	return str
}
