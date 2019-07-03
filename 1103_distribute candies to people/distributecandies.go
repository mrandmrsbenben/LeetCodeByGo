package main

import "fmt"

func main() {
	// candies, num_people := 7, 4
	// candies, num_people := 10, 3
	candies, num_people := 1000000000, 1000
	fmt.Println("Output:", distributeCandies(candies, num_people))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :6.3 MB, 在所有 Go 提交中击败了100.00%的用户
func distributeCandies(candies int, num_people int) []int {
	ans := make([]int, num_people)
	c := 1
	for candies > 0 {
		for i := 0; i < num_people; i++ {
			if candies-c > 0 {
				ans[i] += c
				candies -= c
			} else {
				ans[i] += candies
				candies = 0
				break
			}
			c++
		}
	}
	return ans
}
