package main

import "fmt"

func main() {
	// graph := [][]int{{1, 2}, {3}, {3}, {}}
	graph := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	fmt.Println("Output:", allPathsSourceTarget(graph))
}

//执行用时 :296 ms, 在所有 Go 提交中击败了29.17%的用户
//内存消耗 :34.8 MB, 在所有 Go 提交中击败了100.00%的用户
func allPathsSourceTarget(graph [][]int) [][]int {
	paths := make([][]int, 0)
	var stepTo func(step int, route []int)
	stepTo = func(step int, route []int) {
		for _, n := range graph[step] {
			subroute := make([]int, len(route)+1)
			copy(subroute, route)
			subroute[len(route)] = n
			if n == len(graph)-1 {
				paths = append(paths, subroute)
			} else if n < len(graph) {
				stepTo(n, subroute)
			}
		}
	}
	for _, n := range graph[0] {
		route := make([]int, 2)
		route[0] = 0
		route[1] = n
		if n == len(graph)-1 {
			paths = append(paths, route)
		} else {
			stepTo(n, route)
		}
	}
	return paths
}
