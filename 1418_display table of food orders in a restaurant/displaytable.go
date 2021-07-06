package main

import (
	"fmt"
	"sort"
	"strconv"
)

type TableOrder struct {
	number int
	nostr  string
	food   map[string]int
}

func (t TableOrder) addTableOrder(order []string) {
	t.food[order[2]]++
}

func createTableOrder(order []string) *TableOrder {
	no, _ := strconv.Atoi(order[1])
	food := make(map[string]int)
	food[order[2]] = 1
	return &TableOrder{no, order[1], food}
}

//执行用时：136 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：17.4 MB, 在所有 Go 提交中击败了94.44%的用户
func displayTable(orders [][]string) [][]string {
	foodMap := make(map[string]int)
	tables := make([]*TableOrder, 500)

	var tableNo, tableCnt int
	for _, order := range orders {
		tableNo = getTableNo(order[1])
		if tables[tableNo] == nil {
			tables[tableNo] = createTableOrder(order)
			tableCnt++
		} else {
			tables[tableNo].addTableOrder(order)
		}
		foodMap[order[2]]++
	}

	foodCnt := len(foodMap) + 1
	res := make([][]string, tableCnt+1)
	res[0] = make([]string, foodCnt)
	i := 1
	for f := range foodMap {
		res[0][i] = f
		i++
	}
	sort.Slice(res[0], func(i, j int) bool {
		return res[0][i] < res[0][j]
	})
	res[0][0] = "Table"

	i = 1
	for _, t := range tables {
		if t != nil {
			res[i] = make([]string, foodCnt)
			res[i][0] = t.nostr
			for j := 1; j < foodCnt; j++ {
				res[i][j] = strconv.Itoa(t.food[res[0][j]])
			}
			i++
		}
	}
	return res
}

//执行用时：128 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：17.9 MB, 在所有 Go 提交中击败了83.33%的用户
func displayTable2(orders [][]string) [][]string {
	foodMap := make(map[string]int)
	tables := make(map[string]*TableOrder)

	tableCnt := 0
	for _, order := range orders {
		if t, ok := tables[order[1]]; !ok {
			t = createTableOrder(order)
			tableCnt++
			tables[order[1]] = t
		} else {
			t.addTableOrder(order)
		}
		foodMap[order[2]]++
	}

	ta := make([]*TableOrder, tableCnt)
	i := 0
	for _, t := range tables {
		ta[i] = t
		i++
	}
	sort.Slice(ta, func(i, j int) bool {
		return ta[i].number < ta[j].number
	})

	foodCnt := len(foodMap) + 1
	res := make([][]string, tableCnt+1)
	res[0] = make([]string, foodCnt)
	i = 1
	for f := range foodMap {
		res[0][i] = f
		i++
	}
	sort.Slice(res[0], func(i, j int) bool {
		return res[0][i] < res[0][j]
	})
	res[0][0] = "Table"

	i = 1
	for _, t := range ta {
		if t != nil {
			res[i] = make([]string, foodCnt)
			res[i][0] = t.nostr
			for j := 1; j < foodCnt; j++ {
				res[i][j] = strconv.Itoa(t.food[res[0][j]])
			}
			i++
		}
	}
	return res
}

func getTableNo(tableNumber string) int {
	res, _ := strconv.Atoi(tableNumber)
	return res - 1
}

func main() {
	orders := [][]string{
		// {"Laura", "2", "Bean Burrito"},
		// {"Jhon", "2", "Beef Burrito"},
		// {"Melissa", "2", "Soda"}}
		{"David", "3", "Ceviche"},
		{"Corina", "10", "Beef Burrito"},
		{"David", "3", "Fried Chicken"},
		{"Carla", "5", "Water"},
		{"Carla", "5", "Ceviche"},
		{"Rous", "3", "Ceviche"}}
	// {"James", "12", "Fried Chicken"},
	// {"Ratesh", "12", "Fried Chicken"},
	// {"Amadeus", "12", "Fried Chicken"},
	// {"Adam", "1", "Canadian Waffles"},
	// {"Brianna", "1", "Canadian Waffles"}}
	fmt.Println(displayTable2(orders))
}
