package main

import (
	"fmt"
	"sort"
)

func main() {
	deck := []int{17, 13, 11, 2, 3, 5, 7}
	fmt.Println("Output:", deckRevealedIncreasing(deck))
	// deck := []int{2, 13, 3, 11, 5, 17, 7}
	// fmt.Println("Output:", revealCards(deck))
}

//执行用时 :16 ms, 在所有 Go 提交中击败了95.83%的用户
//内存消耗 :7 MB, 在所有 Go 提交中击败了87.50%的用户
func deckRevealedIncreasing(deck []int) []int {
	if len(deck) == 1 {
		return deck
	}
	cards := make([]int, len(deck))
	sort.Ints(deck)
	cards[len(deck)-1] = deck[len(deck)-1]
	for i := 1; i < len(deck); i++ {
		for j := len(deck) - 1; j > len(deck)-1-i; j-- {
			if j == len(deck)-1 {
				cards[len(deck)-1-i] = cards[j]
			} else {
				cards[j+1] = cards[j]
			}
		}
		cards[len(deck)-1-i], cards[len(deck)-i] = deck[len(deck)-1-i], cards[len(deck)-i-1]
	}
	return cards
}

//执行用时 :44 ms, 在所有 Go 提交中击败了8.33%的用户
//内存消耗 :7.6 MB, 在所有 Go 提交中击败了6.25%的用户
func deckRevealedIncreasingV1(deck []int) []int {
	if len(deck) == 1 {
		return deck
	}
	cards := make([]int, 1)
	sort.Ints(deck)
	cards[0] = deck[len(deck)-1]
	for i := 1; i < len(deck); i++ {
		cards = append([]int{0}, cards...)
		for j := len(cards) - 1; j > 0; j-- {
			if j == len(cards)-1 {
				cards[0] = cards[j]
			} else {
				cards[j+1] = cards[j]
			}
		}
		cards[1], cards[0] = cards[0], deck[len(deck)-1-i]
	}
	return cards
}

// reveal cards simulation
func revealCards(deck []int) []int {
	cards := make([]int, len(deck))
	i := 0
	for len(deck) > 0 {
		fmt.Println(deck)
		cards[i] = deck[0]
		if len(deck) > 2 {
			deck = append(deck[2:], deck[1])
		} else {
			deck = deck[1:]
		}
		i++
	}
	return cards
}
