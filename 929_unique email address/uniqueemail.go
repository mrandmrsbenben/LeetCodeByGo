package main

import (
	"fmt"
	"strings"
)

func main() {
	emails := []string{"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com"}
	fmt.Printf("Number of Unique Emails:%d\n", numUniqueEmails(emails))
}

func numUniqueEmails(emails []string) int {
	em := make(map[string]int)
	var ms []string
	var local string
	for _, s := range emails {
		ms = strings.Split(s, "@")
		if len(ms) != 2 {
			continue
		}
		local = strings.Replace(ms[0], ".", "", -1)
		if strings.Contains(local, "+") {
			index := strings.Index(local, "+")
			local = local[0:index]
		}
		em[local+"@"+ms[1]] = 0
	}
	return len(em)
}
