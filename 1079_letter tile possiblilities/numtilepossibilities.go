package main

import (
	"fmt"
	"strings"
)

func main() {
	tiles := "AAB"
	// tiles := "AAABBC"
	fmt.Println("Output:", numTilePossibilities(tiles))
}

func numTilePossibilities(tiles string) int {
	tmap := make(map[string]int)
	strs := make([]string, len(tiles))
	for i := 0; i < len(tiles); i++ {
		tmap[tiles[i:i+1]] = 1
		strs[i] = tiles[i : i+1]
	}
	bufMap := make(map[string]int)
	l := 1
	for l < len(tiles) {
		for _, s := range strs {
			str := removeFromStr(tiles, s)
			for i := 0; i < len(str); i++ {
				bufMap[s+str[i:i+1]] = 1
			}
		}
		strs = make([]string, len(bufMap))
		i := 0
		for s := range bufMap {
			strs[i] = s
			tmap[s] = 1
			i++
		}
		bufMap = make(map[string]int)
		l++
	}
	return len(tmap)
}

func removeFromStr(s, r string) string {
	var index int
	for i := 0; i < len(r); i++ {
		index = strings.Index(s, r[i:i+1])
		if index != -1 {
			s = s[0:index] + s[index+1:]
		}
	}
	return s
}
