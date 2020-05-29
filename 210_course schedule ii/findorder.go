package main

import "fmt"

func main() {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	fmt.Println("Output: ", findOrder(numCourses, prerequisites))
}

//执行用时 :12 ms, 在所有 Go 提交中击败了87.35%的用户
//内存消耗 :5.9 MB, 在所有 Go 提交中击败了100.00%的用户
func findOrder(numCourses int, prerequisites [][]int) []int {
	ans := make([]int, 0)

	// sub courses
	sc := make([][]int, numCourses)
	for i := range sc {
		sc[i] = make([]int, 0)
	}

	// input degrees
	inds := make([]int, numCourses)
	for _, p := range prerequisites {
		inds[p[0]]++
		sc[p[1]] = append(sc[p[1]], p[0])
	}

	queue := make([]int, 0)
	for i := range inds {
		if inds[i] == 0 {
			queue = append(queue, i)
		}
	}

	buf := make([]int, 0)
	for len(queue) > 0 {
		ans = append(ans, queue...)
		for _, c := range queue {
			for _, n := range sc[c] {
				inds[n]--
				if inds[n] == 0 {
					buf = append(buf, n)
				}
			}
		}
		queue = make([]int, len(buf))
		copy(queue, buf)
		buf = make([]int, 0)
	}
	if len(ans) < numCourses {
		return []int{}
	}
	return ans
}
