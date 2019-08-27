package main

import (
	"fmt"
)

func main() {
	// deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	// target := "0202"
	// deadends := []string{"8888"}
	// target := "0009"
	// deadends := []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	// target := "8888"
	deadends := []string{"0000"}
	target := "8888"
	fmt.Println(openLock(deadends, target))
}

//执行用时 :116 ms, 在所有 Go 提交中击败了60.00%的用户
//内存消耗 :7.3 MB, 在所有 Go 提交中击败了42.86%的用户
func openLock(deadends []string, target string) int {
	turns := 0
	deadendsMap := make(map[string]int)
	//add all deadends to map
	for _, k := range deadends {
		deadendsMap[k] = 1
	}
	if deadendsMap["0000"] == 1 {
		return -1
	} else if target == "0000" {
		return 0
	}
	//array for every round state
	states := make([]string, 1)
	var stateF, stateB string
	var l int
	states[0] = "0000"
	for len(states) > 0 {
		turns++
		l = len(states)
		for i := 0; i < l; i++ {
			for j := 0; j < 4; j++ {
				stateF, stateB = turnOnce(states[i], j)
				if stateF == target || stateB == target {
					return turns
				}
				if _, ok := deadendsMap[stateF]; !ok {
					deadendsMap[stateF] = 1
					states = append(states, stateF)
				}
				if _, ok := deadendsMap[stateB]; !ok {
					deadendsMap[stateB] = 1
					states = append(states, stateB)
				}
			}
		}
		states = states[l:]
	}
	return -1
}

func turnOnce(state string, bit int) (string, string) {
	nums := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var s1, s2 string
	n := state[bit] - '0'
	s1 = nums[(n+1)%10]
	s2 = nums[(n+9)%10]
	switch bit {
	case 0:
		return s1 + state[1:], s2 + state[1:]
	case 3:
		return state[0:3] + s1, state[0:3] + s2
	default:
		return state[0:bit] + s1 + state[bit+1:], state[0:bit] + s2 + state[bit+1:]
	}
}
