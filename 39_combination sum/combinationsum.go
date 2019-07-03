package main

import "fmt"

func main() {
	// candidates := []int{2, 3, 6, 7}
	// target := 7
	candidates := []int{2, 3, 5}
	target := 8
	fmt.Println("Output:", combinationSum(candidates, target))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了96.76%的用户
//内存消耗 :4.6 MB, 在所有 Go 提交中击败了36.36%的用户
func combinationSum(candidates []int, target int) [][]int {
	comb := [][]int{}
	for i := range candidates {
		if candidates[i] == target {
			comb = append(comb, []int{candidates[i]})
		} else if candidates[i] < target {
			rtn := combinationSum(candidates[i:], target-candidates[i])
			for j := range rtn {
				if len(rtn[j]) == 0 {
					continue
				}
				comb = append(comb, append([]int{candidates[i]}, rtn[j]...))
			}
		}
	}
	return comb
}
