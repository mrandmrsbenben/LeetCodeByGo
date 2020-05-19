package main

import "fmt"

//执行用时 :204 ms, 在所有 Go 提交中击败了71.21%的用户
//内存消耗 :8.4 MB, 在所有 Go 提交中击败了100.00%的用户
// StockSpanner stock spanner
type StockSpanner struct {
	prices []int
	spans  []int
}

// Constructor create new stock spanner
func Constructor() StockSpanner {
	return StockSpanner{[]int{}, []int{}}
}

// Next get the span of current price
func (this *StockSpanner) Next(price int) int {
	span := 1
	i := len(this.spans) - 1
	for i >= 0 {
		if this.prices[i] > price {
			break
		}
		span += this.spans[i]
		i -= this.spans[i]
	}
	this.prices = append(this.prices, price)
	this.spans = append(this.spans, span)
	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
func main() {
	obj := Constructor()
	vals := []int{100, 80, 60, 70, 60, 75, 85}
	for _, v := range vals {
		fmt.Println("Input:", v, "Output ", obj.Next(v))
	}
}
