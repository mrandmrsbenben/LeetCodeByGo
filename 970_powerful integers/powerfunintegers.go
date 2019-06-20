package main

import "fmt"

func main() {
	// x, y := 2, 3
	// bound := 10
	x, y := 1, 1
	bound := 0
	fmt.Printf("Output: %v\n", powerfulIntegers(x, y, bound))
}

// 执行用时 :0 ms, 在所有Go提交中击败了100.00%的用户
// 内存消耗 :2 MB, 在所有Go提交中击败了45.45%的用户
func powerfulIntegers(x int, y int, bound int) []int {
	pi := make([]int, 0)
	pim := make(map[int]int)
	px := pows(x, bound)
	py := pows(y, bound)
	for _, m := range px {
		for _, n := range py {
			if m+n <= bound {
				if _, ok := pim[m+n]; !ok {
					pi = append(pi, m+n)
					pim[m+n] = 1
				}
			} else {
				break
			}
		}
		if len(pi) == 0 || pi[len(pi)-1] > bound {
			break
		}
	}
	return pi
}

func pows(n, bound int) []int {
	p := []int{1}
	if n == 1 {
		return p
	}
	for {
		if n*p[len(p)-1] <= bound {
			p = append(p, n*p[len(p)-1])
		} else {
			break
		}
	}
	return p
}
