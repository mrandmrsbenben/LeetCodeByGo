package main

import "fmt"

func main() {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	// list1 := []string{"Shogun", "KFC", "Tapioca Express", "Burger King"}
	// list2 := []string{"KFC", "Shogun", "Burger King"}
	fmt.Printf("Output: %v\n", findRestaurant(list1, list2))
}

func findRestaurant(list1 []string, list2 []string) []string {
	common := make([]string, 0)
	min := len(list1) + len(list2)
	if len(list2) < len(list1) {
		list1, list2 = list2, list1
	}
	for i := range list2 {
		for j := range list1 {
			if list2[i] == list1[j] {
				if i+j < min {
					common = []string{list1[j]}
					min = i + j
				} else if i+j == min {
					common = append(common, list1[j])
				}
				break
			}
		}
	}
	return common
}
