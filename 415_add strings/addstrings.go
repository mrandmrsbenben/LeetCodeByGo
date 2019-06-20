package main

import (
	"fmt"
)

func main() {
	num1 := "1234"
	num2 := "97890"
	// for i := 0; i < 5100; i++ {
	// 	num2 = num2 + "9"
	// }
	fmt.Printf("Output: %v\n", addStrings(num1, num2))
}

func addStrings(num1 string, num2 string) string {
	im := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4,
		"5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	sa := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	sum := make([]string, 0)
	addup, bitsum := 0, 0
	j := len(num2) - 1
	for i := len(num1) - 1; i >= 0; i-- {
		if j >= 0 {
			bitsum = im[num1[i:i+1]] + im[num2[j:j+1]] + addup
			if bitsum > 9 {
				addup = bitsum / 10
			} else {
				addup = 0
			}
			sum = append(sum, sa[bitsum%10])
			j = j - 1
		} else {
			bitsum = im[num1[i:i+1]] + addup
			if bitsum > 9 {
				addup = bitsum / 10
			} else {
				addup = 0
			}
			sum = append(sum, sa[bitsum%10])
		}
	}
	for {
		if addup == 0 {
			break
		}
		sum = append(sum, sa[addup%10])
		addup = addup / 10
	}
	sums := ""
	for i := len(sum) - 1; i >= 0; i-- {
		sums = sums + sum[i]
	}
	return sums
}
