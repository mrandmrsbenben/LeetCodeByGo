package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(strings.Split(s, ""))
	fmt.Printf("Output: %v\n", isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ss := strings.Split(s, "")
	ts := strings.Split(t, "")
	sort.Strings(ss)
	sort.Strings(ts)
	return strings.Join(ss, "") == strings.Join(ts, "")
}
