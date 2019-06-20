package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := 0
	output := findComplement(input)
	fmt.Printf("Input:%d %b, Output:%d %b\n", input, input, output, output)
}

func findComplement(num int) int {
	in := fmt.Sprintf("%b", num)
	var out string
	for i := 0; i < len(in); i++ {
		b, _ := strconv.Atoi(string(in[i]))
		out = out + strconv.Itoa((b+1)%2)
	}
	comp, _ := strconv.ParseInt(out, 2, 32)
	return int(comp)
}
