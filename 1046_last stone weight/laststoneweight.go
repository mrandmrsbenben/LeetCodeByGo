package main

import (
	"fmt"
	"sort"
)

func main() {
	stones := []int{316, 157, 73, 106, 771, 828, 46, 212, 926, 604, 600, 992, 71, 51, 477, 869, 425, 405, 859, 924, 45, 187, 283, 590, 303, 66, 508, 982, 464, 398}
	fmt.Printf("Output: %d\n", lastStoneWeight(stones))
}

func lastStoneWeight(stones []int) int {
	if len(stones) == 2 {
		if stones[1]-stones[0] < 0 {
			return stones[0] - stones[1]
		}
		return stones[1] - stones[0]
	} else if len(stones) == 1 {
		return stones[0]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(stones)))
	fmt.Println(stones)
	smashed := append(stones[2:], stones[0]-stones[1])
	// for i := 0; i < len(stones); i = i + 2 {
	// 	if i+1 < len(stones) {
	// 		smashed = append(smashed, stones[i]-stones[i+1])
	// 	} else {
	// 		smashed = append(smashed, stones[i])
	// 	}
	// }
	return lastStoneWeight(smashed)
}
