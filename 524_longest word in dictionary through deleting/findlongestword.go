package main

import "fmt"

func main() {
	s := "abpcplea"
	d := []string{"ale","apple","monkey","plea"}
	fmt.Println("Output: ", findLongestWord(s,d))
}

//执行用时：16 ms, 在所有 Go 提交中击败了90.57%的用户
//内存消耗：7.7 MB, 在所有 Go 提交中击败了93.40%的用户
func findLongestWord(s string, d []string) string {
	res := ""
	
	for _,str := range d{
		if len(str) > len(res){
			i,j := 0, 0
			for i < len(str) && j < len(s){
				if str[i] == s[j]{
					i++
					j++
				}else{
					j++
				}
			}
			if i == len(str){
                if len(str) > len(res){
                    res = str
                }else if res > str{
                    res = str
                }
			}
		}
	}

	return res
}