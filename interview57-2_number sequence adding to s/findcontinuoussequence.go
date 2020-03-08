package main

import "fmt"

// https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
func main() {
	target := 10000
	fmt.Println("Output: ", findContinuousSequence(target))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :6 MB, 在所有 Go 提交中击败了100.00%的用户
func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	seq := make([]int, 0)
	buf := make([]int, 0)
	sum := 0
	for i := 1; i <= (target+1)/2; i++ {
		sum += i
		seq = append(seq, i)
		if sum > target {
			for sum > target {
				sum -= seq[0]
				seq = seq[1:]
			}
		}
		if sum == target {
			buf = make([]int, len(seq))
			copy(buf, seq)
			res = append(res, buf)
			sum -= seq[0]
			seq = seq[1:]
		}
	}
	return res
}

//执行用时 :376 ms, 在所有 Go 提交中击败了5.63%的用户
//内存消耗 :6.8 MB, 在所有 Go 提交中击败了100.00%的用户
func findContinuousSequenceV2(target int) [][]int {
	res := make([][]int, 0)
	for i := 1; i <= (target+1)/2; i++ {
		seq := make([]int, 0)
		sum := 0
		for j := (target+1)/2 - i + 1; sum < target && j > 0; j-- {
			sum += j
			seq = append([]int{j}, seq...)
		}
		if len(seq) > 1 && sum == target {
			res = append([][]int{seq}, res...)
		}
	}
	return res
}

//执行用时 :48 ms, 在所有 Go 提交中击败了11.27%的用户
//内存消耗 :6.6 MB, 在所有 Go 提交中击败了100.00%的用户
func findContinuousSequenceV1(target int) [][]int {
	res := make([][]int, 0)
	for i := 1; i <= (target+1)/2; i++ {
		seq := make([]int, 0)
		sum := 0
		for j := i; sum < target; j++ {
			sum += j
			seq = append(seq, j)
		}
		if len(seq) > 1 && sum == target {
			res = append(res, seq)
		}
	}

	return res
}
