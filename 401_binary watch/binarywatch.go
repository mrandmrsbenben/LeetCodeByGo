package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := 8
	fmt.Printf("Output: %v\n", readBinaryWatch(num))
}

func readBinaryWatch(num int) []string {
	if num == 0 {
		return []string{"0:00"}
	}
	str := make([]string, 0)
	if num > 8 || num < 0 {
		return str
	}
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if strings.Count(strconv.FormatInt(int64(i), 2), "1")+
				strings.Count(strconv.FormatInt(int64(j), 2), "1") == num {
				if j < 10 {
					str = append(str, strconv.Itoa(i)+":0"+strconv.Itoa(j))
				} else {
					str = append(str, strconv.Itoa(i)+":"+strconv.Itoa(j))
				}
			}
		}
	}
	return str
}
