package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// common.MakeTestCases()
	str1 := []string{"9001 discuss.leetcode.com"}
	fmt.Printf("Output: %v\n", subdomainVisits(str1))
	fmt.Printf("Expect: {\"9001 discuss.leetcode.com\", \"9001 leetcode.com\", \"9001 com\"}\n")
	str2 := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	fmt.Printf("Output: %v\n", subdomainVisits(str2))
	fmt.Printf("Expect: {\"901 mail.com\",\"50 yahoo.com\",\"900 google.mail.com\",\"5 wiki.org\",\"5 org\",\"1 intel.mail.com\",\"951 com\"}\n")
}

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	var d string
	var f, ds []string
	var cnt int
	for _, s := range cpdomains {
		f = strings.Fields(s)
		cnt, _ = strconv.Atoi(f[0])
		ds = strings.Split(f[1], ".")
		d = ""
		for i := len(ds) - 1; i >= 0; i-- {
			d = ds[i] + d
			if p, ok := m[d]; ok {
				m[d] = p + cnt
			} else {
				m[d] = cnt
			}
			d = "." + d
		}
	}
	vs := make([]string, 0)
	for k, v := range m {
		vs = append(vs, strconv.Itoa(v)+" "+k)
	}
	return vs
}
