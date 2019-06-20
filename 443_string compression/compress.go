package main

import (
	"fmt"
	"strconv"
)

func main() {
	// chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	// chars := []byte{'a'}
	chars := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	fmt.Println("Output:", compress(chars))
}

//执行用时 :16 ms, 在所有Go提交中击败了91.53%的用户
//内存消耗 :6.7 MB, 在所有Go提交中击败了37.50%的用户
func compress(chars []byte) int {
	if len(chars) <= 1 {
		return len(chars)
	}
	cnt := 1
	c := chars[0]
	comp := make([]byte, 0)
	for i := 1; i < len(chars); i++ {
		if chars[i] != c {
			comp = append(comp, c)
			if cnt > 1 {
				comp = append(comp, []byte(strconv.Itoa(cnt))...)
			}
			cnt = 1
			c = chars[i]
		} else {
			cnt++
		}
	}
	comp = append(comp, c)
	if cnt > 1 {
		comp = append(comp, []byte(strconv.Itoa(cnt))...)
	}
	if len(comp) > len(chars) {
		for i := 0; i < len(chars); i++ {
			chars[i] = comp[i]
		}
		chars = append(chars, comp[len(chars):]...)
	} else {
		for i := 0; i < len(comp); i++ {
			chars[i] = comp[i]
		}
		chars = chars[0:len(comp)]
	}
	return len(chars)
}
