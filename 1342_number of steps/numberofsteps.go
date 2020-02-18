package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Output: ", numberOfSteps(8))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func numberOfSteps(num int) int {
	binstr := strconv.FormatInt(int64(num), 2)
	steps := 0
	if binstr[len(binstr)-1] == '1' {
		steps++
	}
	for i := 0; i < len(binstr)-1; i++ {
		if binstr[i] == '1' {
			steps += 2
		} else {
			steps++
		}
	}
	return steps
}
