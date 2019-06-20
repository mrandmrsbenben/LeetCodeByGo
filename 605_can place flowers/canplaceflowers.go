package main

import "fmt"

func main() {
	// flowerbed := []int{1, 0, 0, 0, 0, 0, 1}
	// n := 2
	flowerbed := []int{0, 0, 1, 0, 0}
	n := 3
	fmt.Println("Output:", canPlaceFlowers(flowerbed, n))
}

//执行用时 :24 ms, 在所有 Go 提交中击败了97.87%的用户
//内存消耗 :5.9 MB, 在所有 Go 提交中击败了22.73%的用户
func canPlaceFlowers(flowerbed []int, n int) bool {
	plots := 0
	cnt := 1
	for _, f := range flowerbed {
		if f == 1 {
			plots += getFlowerPlots(cnt)
			cnt = 0
		} else {
			cnt++
		}
	}
	if cnt > 0 {
		plots += getFlowerPlots(cnt + 1)
	}
	return plots >= n
}

func getFlowerPlots(n int) int {
	if n < 3 {
		return 0
	}
	if n%2 == 0 {
		return n/2 - 1
	}
	return n / 2
}
