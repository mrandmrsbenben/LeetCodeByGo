package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// secret := "1807"
	// guess := "7810"
	// secret := "1123"
	// guess := "0111"
	// secret := "11"
	// guess := "11"
	secret := "11"
	guess := "10"
	fmt.Println("Output:", getHint(secret, guess))
}

func getHint(secret string, guess string) string {
	a, b := 0, 0
	var index int
	for i := range guess {
		if secret[i] == guess[i] {
			a++
			secret = markStr(secret, i)
			guess = markStr(guess, i)
		}
	}
	for i := range guess {
		if guess[i] == '#' {
			continue
		}
		index = strings.Index(secret, guess[i:i+1])
		if index != -1 {
			b++
			secret = markStr(secret, index)
		}
	}
	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}

func markStr(str string, index int) string {
	switch index {
	case 0:
		str = "#" + str[1:]
	case len(str) - 1:
		str = str[0:len(str)-1] + "#"
	default:
		str = str[0:index] + "#" + str[index+1:]
	}
	return str
}
