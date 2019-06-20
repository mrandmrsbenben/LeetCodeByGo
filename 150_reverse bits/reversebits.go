package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var num uint32 = 43261596
	fmt.Printf("Output: %d\n", reverseBits(num))
}

func reverseBits(num uint32) uint32 {
	str := strconv.FormatUint(uint64(num), 2)
	rev := ""
	for i := len(str) - 1; i >= 0; i-- {
		rev = rev + str[i:i+1]
	}
	if len(rev) < 32 {
		rev = rev + strings.Repeat("0", 32-len(rev))
	}
	n, _ := strconv.ParseUint(rev, 2, 64)
	return uint32(n)
}
