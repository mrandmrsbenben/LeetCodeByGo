package main

import "fmt"

func main() {
	str1 := "up"
	b1 := []byte(str1)
	str2 := "Hannah"
	b2 := []byte(str2)
	reverseString(b1)
	reverseString(b2)
	fmt.Printf("Input:%v, Output:%v\n", []byte(str1), string(b1))
	fmt.Printf("Input:%v, Output:%v\n", []byte(str2), string(b2))
}

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
