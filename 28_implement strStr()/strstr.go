package main

import (
	"fmt"
	"strings"
)

func main() {
	haystack := "hello"
	needle := "le"
	fmt.Printf("Output: %d\n", strStr(haystack, needle))
}

// 执行用时 : 0 ms, 在Implement strStr()的Go提交中击败了100.00% 的用户
// 内存消耗 : 2.4 MB, 在Implement strStr()的Go提交中击败了35.40% 的用户
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	} else if haystack == "" {
		return -1
	}
	ln := len(needle)
	for i := range haystack {
		if i+ln > len(haystack) {
			break
		}
		if haystack[i:i+ln] == needle {
			return i
		}
	}
	return -1
}

func strStrV1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
