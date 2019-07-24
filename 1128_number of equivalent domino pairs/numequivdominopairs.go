package main

import (
	"fmt"
	"strconv"
)

func main() {
	// dominoes := [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}
	dominoes := [][]int{{1, 2}, {2, 1}, {1, 2}, {2, 1}}
	fmt.Println("Output:", numEquivDominoPairs(dominoes))
}

//执行用时 :44 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :7.6 MB, 在所有 Go 提交中击败了100.00%的用户
func numEquivDominoPairs(dominoes [][]int) int {
	dm := make(map[string]int)
	var key string
	for _, d := range dominoes {
		key = getDominoKey(d)
		if c, ok := dm[key]; ok {
			dm[key] = c + 1
		} else {
			dm[key] = 1
		}
	}
	cnt := 0
	for _, v := range dm {
		cnt += v * (v - 1) / 2
	}
	return cnt
}

func getDominoKey(d []int) string {
	if d[0] <= d[1] {
		return strconv.Itoa(d[0]) + "," + strconv.Itoa(d[1])
	}
	return strconv.Itoa(d[1]) + "," + strconv.Itoa(d[0])
}

func numEquivDominoPairsV1(dominoes [][]int) int {
	cnt := 0
	for i := range dominoes {
		for j := i + 1; j < len(dominoes); j++ {
			if isEqual(dominoes[i], dominoes[j]) {
				cnt++
				break
			}
		}
	}
	return cnt
}

func isEqual(a, b []int) bool {
	if a[0] == b[0] && a[1] == b[1] {
		return true
	} else if a[0] == b[1] && a[1] == b[0] {
		return true
	}
	return false
}
