package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	num := -1
	fmt.Println(strconv.FormatInt(int64(num), 16))
	fmt.Printf("Output: %v\n", toHex(num))
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	hm := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	if num < 0 {
		num = int(math.Pow(2, 32)) + num
	}
	bits := int(math.Log(float64(num)) / math.Log(16))
	hex := ""
	for i := bits; i >= 0; i-- {
		if i > 0 {
			b := int(math.Pow(16, float64(i)))
			hex = hex + hm[num/b]
			num = num % b
		} else {
			hex = hex + hm[num]
		}
	}
	return hex
}
