package main

import "fmt"

func main() {
	// s := "10#11#12"
	// s := "1326#"
	s := "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"
	fmt.Println("Output: ", freqAlphabets(s))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.6 MB, 在所有 Go 提交中击败了16.67%的用户
func freqAlphabets(s string) string {
	result := ""
	m := make(map[string]string)
	keys := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10#", "11#", "12#", "13#",
		"14#", "15#", "16#", "17#", "18#", "19#", "20#", "21#", "22#", "23#", "24#", "25#", "26#"}
	alphabets := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for i := range keys {
		m[keys[i]] = alphabets[i]
	}

	i := len(s) - 1
	for i >= 0 {
		if s[i] == '#' {
			result = m[s[i-2:i+1]] + result
			i -= 2
		} else {
			result = m[s[i:i+1]] + result
		}
		i--
	}
	return result
}
