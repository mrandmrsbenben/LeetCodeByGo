package main

import (
	"fmt"
	"strings"
)

type jas struct {
	jewels string
	stones string
}

func main() {
	testcase := []jas{{"aA", "aAAbbbba"}, {"zZ", "Z"}}
	for _, t := range testcase {
		fmt.Printf("Jewels:%s,Stones:%s,Nums Of Jewels:%d\n",
			t.jewels, t.stones, numJewelsInStones(t.jewels, t.stones))
	}
}

func numJewelsInStones(J string, S string) int {
	num := 0
	for i := 0; i < len(J); i++ {
		num = num + strings.Count(S, string(J[i]))
	}
	return num
}
