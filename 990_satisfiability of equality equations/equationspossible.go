package main

import "fmt"

func main() {
	// 		a,b,c,d,e,f,g,
	// a	1,1,1,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
	// b	1,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
	// c	1,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
	// d	0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
	// e	0,1,0,0,1,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
	// f	1,0,1,0,2,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
	// g	0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0

	// equations := []string{"a==b", "b!=d", "d==a"}
	// equations := []string{"a==b", "b!=a"}
	// equations := []string{"a==b", "b==c", "a==c"}
	// equations := []string{"c==c", "b==d", "f!=g"}
	// equations := []string{"c==c", "f!=a", "f==b", "b==c"}
	// equations := []string{"a==b", "e==c", "b==c", "a!=e"}
	equations := []string{"f==a", "a==b", "f!=e", "a==c", "b==e", "c==f"}
	fmt.Println("Output: ", equationsPossible(equations))
}

func equationsPossible(equations []string) bool {
	vals := [26][26]int{}
	for i := range vals {
		vals[i][i] = 1
	}

	for _, e := range equations {
		if e[1] == '=' {
			if vals[e[0]-'a'][e[3]-'a'] == 2 || vals[e[3]-'a'][e[0]-'a'] == 2 {
				return false
			}
			vals[e[0]-'a'][e[3]-'a'] = 1
			vals[e[3]-'a'][e[0]-'a'] = 1
		} else {
			if vals[e[0]-'a'][e[3]-'a'] == 1 || vals[e[3]-'a'][e[0]-'a'] == 1 {
				return false
			}
			vals[e[0]-'a'][e[3]-'a'] = 2
			vals[e[3]-'a'][e[0]-'a'] = 2
		}
	}

	for i := range vals {
		for j := 0; j < 26; j++ {
			if j != i && vals[i][j] == 1 {
				for k := 0; k < 26; k++ {
					if vals[j][k] == 1 {
						if vals[i][k] == 2 {
							return false
						}
						vals[i][k] = 1
					}
				}
			}
		}
	}

	return true
}
