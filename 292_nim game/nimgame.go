package main

import "fmt"

func main() {
	fmt.Println(canWinNim(1220))
}

func canWinNim(n int) bool {
	return n%4 != 0
}
