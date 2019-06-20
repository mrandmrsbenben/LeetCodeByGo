package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
	b := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	fmt.Printf("Output: %v\n", addBinary(a, b))
}

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	b = strings.Repeat("0", len(a)-len(b)) + b
	c := ""
	cnt := 0
	for i := len(a) - 1; i >= 0; i-- {
		if a[i:i+1] == "1" && b[i:i+1] == "1" {
			if cnt > 0 {
				c = "1" + c
			} else {
				c = "0" + c
				cnt = cnt + 1
			}
		}
		if (a[i:i+1] == "1" && b[i:i+1] == "0") ||
			(a[i:i+1] == "0" && b[i:i+1] == "1") {
			if cnt > 0 {
				c = "0" + c
			} else {
				c = "1" + c
			}
		}
		if a[i:i+1] == "0" && b[i:i+1] == "0" {
			if cnt > 0 {
				c = "1" + c
				cnt = cnt - 1
			} else {
				c = "0" + c
			}
		}
	}
	if cnt > 0 {
		c = "1" + c
	}
	return c
}
