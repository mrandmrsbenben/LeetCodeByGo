package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println("Output:", groupAnagrams(strs))
}

//执行用时 :452 ms, 在所有 Go 提交中击败了84.23%的用户
//内存消耗 :87.9 MB, 在所有 Go 提交中击败了27.78%的用户
func groupAnagrams(strs []string) [][]string {
	var key string
	gm := make(map[string][]string)
	for _, str := range strs {
		key = getKey(str)
		if s, ok := gm[key]; ok {
			s = append(s, str)
			gm[key] = s
		} else {
			s = make([]string, 1)
			s[0] = str
			gm[key] = s
		}
	}
	gs := make([][]string, len(gm))
	i := 0
	for _, s := range gm {
		gs[i] = s
		i++
	}
	return gs
}

func getKey(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
