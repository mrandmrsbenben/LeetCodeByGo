package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println("Output: ", validIPAddress("172.16.254.-1"))
	// fmt.Println("Output: ", validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334"))
	// fmt.Println("Output: ", validIPAddress("256.256.256.256"))
	fmt.Println("Output: ", validIPAddress("20EE:Fb8:85a3:0:0:8A2E:0370:7334"))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户
func validIPAddress(IP string) string {
	res := []string{"Neither", "IPv4", "IPv6"}
	if strings.Count(IP, ".") == 3 {
		strs := strings.Split(IP, ".")
		for _, s := range strs {
			if !validIPv4(s) {
				return res[0]
			}
		}
		return res[1]
	} else if strings.Count(IP, ":") == 7 {
		strs := strings.Split(IP, ":")
		for _, s := range strs {
			if !validIPv6(s) {
				return res[0]
			}
		}
		return res[2]
	}
	return res[0]
}

func validIPv4(s string) bool {
	val, err := strconv.Atoi(s)
	if err != nil {
		return false
	} else if val < 0 || val > 255 {
		return false
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' || s[i] == '-' || s[i] == '+' {
			return false
		}
	}

	return true
}

func validIPv6(s string) bool {
	if len(s) > 4 || len(s) == 0 {
		return false
	}
	for _, c := range s {
		if (c >= '0' && c <= '9') ||
			(c >= 'a' && c <= 'f') ||
			(c >= 'A' && c <= 'F') {
			continue
		}
		return false
	}
	return true
}
