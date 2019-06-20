package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	logs := []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}
	fmt.Printf("Output: %v\n", reorderLogFiles(logs))
}

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i int, j int) bool {
		si := strings.SplitN(logs[i], " ", 2)
		sj := strings.SplitN(logs[j], " ", 2)
		if strings.Contains("01234567890", si[1][0:1]) &&
			strings.Contains("01234567890", sj[1][0:1]) {
			return false
		} else if strings.Contains("01234567890", si[1][0:1]) {
			return false
		} else if strings.Contains("01234567890", sj[1][0:1]) {
			return true
		}
		if si[1] == sj[1] {
			return si[0] < sj[0]
		}
		return si[1] < sj[1]
	})
	return logs
}
