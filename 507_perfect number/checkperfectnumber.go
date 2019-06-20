package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Output:", checkPerfectNumber(56))
}

//执行用时 :0 ms, 在所有Go提交中击败了100.00%的用户
//内存消耗 :2 MB, 在所有Go提交中击败了50.00%的用户
func checkPerfectNumber(num int) bool {
	if num == 0 {
		return false
	}
	sum := 0
	sn := int(math.Sqrt(float64(num)))
	for i := sn; i >= 1; i-- {
		if num%i == 0 {
			if num/i > i {
				sum += num / i
			}
			sum += i
		}
	}
	return sum == num*2
}

//执行用时 :1584 ms, 在所有Go提交中击败了23.53%的用户
//内存消耗 :2 MB, 在所有Go提交中击败了50.00%的用户
func checkPerfectNumberV0(num int) bool {
	if num == 0 {
		return false
	}
	sum := 0
	factors := make(map[int]int)
	var findFactors func(base, num int)
	findFactors = func(base, num int) {
		if _, ok := factors[num]; ok {
			return
		}
		factors[num] = 1
		sum += num
		if num%2 == 0 {
			n := num / 2
			findFactors(base, n)
			findFactors(base, base/n)
		} else {
			for i := num / 3; i >= 1; i -= 2 {
				if num%i == 0 {
					findFactors(base, i)
					findFactors(base, base/i)
					findFactors(base, num/i)
				}
			}
		}
	}
	findFactors(num, num)
	return sum == num*2
}

// func findFactors(base, num int, factors map[int]int) {
// 	if _, ok := factors[num]; ok {
// 		return
// 	}
// 	factors[num] = 1
// 	if num%2 == 0 {
// 		n := num / 2
// 		findFactors(base, 2, factors)
// 		findFactors(base, n, factors)
// 		findFactors(base, base/n, factors)
// 	} else {
// 		for i := num / 3; i >= 1; i -= 2 {
// 			if num%i == 0 {
// 				findFactors(base, i, factors)
// 				findFactors(base, base/i, factors)
// 				findFactors(base, num/i, factors)
// 			}
// 		}
// 	}
// }
