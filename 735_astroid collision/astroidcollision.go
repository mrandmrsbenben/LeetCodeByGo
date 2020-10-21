package main

import "fmt"

func main() {
	// asteroids := []int{5, 10, -5}
	// asteroids := []int{10, 2, -5, -5}
	// asteroids := []int{-2, -1, 1, 2}
	asteroids := []int{1, -2, -2, -2}
	fmt.Println("Output: ", asteroidCollision(asteroids))
}

//执行用时：8 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：4.5 MB, 在所有 Go 提交中击败了88.64%的用户
func asteroidCollision(asteroids []int) []int {
	res := make([]int, 0)
	for _, a := range asteroids {
		if a > 0 || len(res) == 0 {
			res = append(res, a)
		} else {
			for j := len(res) - 1; j >= 0; j-- {
				if res[j] < 0 {
					res = append(res[:j+1], a)
					break
				} else if res[j] > -1*a {
					res = res[:j+1]
					break
				} else if res[j] == -1*a {
					res = res[:j]
					break
				} else if j == 0 {
					res = []int{a}
				}
			}
		}
	}
	return res
}
