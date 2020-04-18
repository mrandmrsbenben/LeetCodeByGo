package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("Output: ", maxArea(height))
}

//执行用时 :16 ms, 在所有 Go 提交中击败了80.47%的用户
//内存消耗 :5.8 MB, 在所有 Go 提交中击败了16.67%的用户
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	max, cur := 0, 0
	for l < r {
		cur = area(height[l], height[r], r-l)
		if max < cur {
			max = cur
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}

func area(l, r int, dis int) int {
	if l > r {
		return r * dis
	}
	return l * dis
}
