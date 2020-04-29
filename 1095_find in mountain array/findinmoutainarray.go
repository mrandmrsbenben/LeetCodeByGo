package main

import "fmt"

//MountainArray Mountain Array
type MountainArray struct {
	array   []int
	callCnt int
}

func createMA(array []int) *MountainArray {
	return &MountainArray{array, 0}
}

func (t *MountainArray) get(index int) int {
	t.callCnt++
	fmt.Println("Get :", t.callCnt)
	return t.array[index]
}

func (t *MountainArray) length() int {
	return len(t.array)
}

func main() {
	// array := []int{1, 2, 2, 4, 5, 3, 1}
	// array := []int{1, 2, 4, 5, 3, 2, 1}
	// array := []int{1, 5, 3, 2}
	// array := []int{0, 1, 2, 4, 2, 1}
	// target := 2
	// array := []int{1, 2, 5, 3}
	// target := 5
	// array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	// target := 21
	array := []int{1, 6, 11, 16, 21, 26, 31, 36, 41, 46, 51, 56, 61, 66, 71, 76, 81, 86, 91, 96, 101, 106, 111, 116, 121, 126, 131, 136, 141, 146, 151, 156, 161, 166, 171, 176, 181, 186, 191, 196, 201, 206, 211, 216, 221, 226, 231, 236, 241, 246, 251, 256, 261, 266, 271, 276, 281, 286, 291, 296, 301, 306, 311, 316, 321, 326, 331, 336, 341, 346, 351, 356, 361, 366, 371, 376, 381, 386, 391, 396, 401, 406, 411, 416, 421, 426, 431, 436, 441, 446, 451, 456, 461, 466, 471, 476, 481, 486, 491, 496, 501, 496, 491, 486, 481, 476, 471, 466, 461, 456, 451, 446, 441, 436, 431, 426, 421, 416, 411, 406}
	target := 481
	fmt.Println("Output: ", findInMountainArray(target, createMA(array)))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :3.6 MB, 在所有 Go 提交中击败了100.00%的用户
func findInMountainArray(target int, mountainArr *MountainArray) int {
	array := make([]int, mountainArr.length())
	for i := range array {
		array[i] = -1
	}

	peak := findPeak(array, mountainArr)
	if getVal(peak, array, mountainArr) == target {
		return peak
	}

	l, r := 0, peak
	var i, val int
	for l < r {
		i = l + (r-l)/2
		val = getVal(i, array, mountainArr)
		if val == target {
			return i
		} else if l == i {
			break
		}

		if val > target {
			r = i
		} else {
			l = i
		}
	}

	l, r = peak, mountainArr.length()-1
	for l < r {
		i = l + (r-l)/2
		val = getVal(i, array, mountainArr)
		if val == target {
			return i
		} else if l == i {
			if getVal(r, array, mountainArr) == target {
				return r
			}
			break
		}

		if val > target {
			l = i
		} else {
			r = i
		}
	}
	return -1
}

func findPeak(array []int, mountainArr *MountainArray) int {
	l, r := 0, mountainArr.length()-1
	var i int
	var mid, next int
	for l < r {
		i = l + (r-l)/2
		if l == i {
			break
		}
		mid = getVal(i, array, mountainArr)
		next = getVal(i+1, array, mountainArr)
		if mid < next {
			l = i
		} else {
			r = i
		}
	}
	return i
}

func getVal(index int, array []int, ma *MountainArray) int {
	if array[index] == -1 {
		array[index] = ma.get(index)
	}
	return array[index]
}
