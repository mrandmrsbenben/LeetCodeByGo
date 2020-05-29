package main

import "fmt"

func main() {
	numCourses := 6
	prerequisites := [][]int{{3, 0}, {3, 1}, {4, 1}, {4, 2}, {5, 3}, {5, 4}}
	// numCourses := 2
	// prerequisites := [][]int{{1, 0}, {0, 1}}
	fmt.Println("Output: ", canFinish(numCourses, prerequisites))
}

//执行用时 :8 ms, 在所有 Go 提交中击败了99.43%的用户
//内存消耗 :5.8 MB, 在所有 Go 提交中击败了100.00%的用户
func canFinish(numCourses int, prerequisites [][]int) bool {
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
	// taken course count
	cnt := len(queue)

	buf := make([]int, 0)
	for len(queue) > 0 {
		for _, c := range queue {
			for _, n := range sc[c] {
				inds[n]--
				if inds[n] == 0 {
					buf = append(buf, n)
					cnt++
				}
			}
		}
		queue = make([]int, len(buf))
		copy(queue, buf)
		buf = make([]int, 0)
	}

	return cnt == numCourses
}
