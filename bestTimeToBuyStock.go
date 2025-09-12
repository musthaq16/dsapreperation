package main

import "fmt"

func bestTimeToBuyStock(stocks []int) int {
	maxProfit := 0
	buy := stocks[0]
	for i := 1; i < len(stocks); i++ {
		if buy > stocks[i] {
			buy = stocks[i]
		} else {
			prof := stocks[i] - buy
			maxProfit = max(prof, maxProfit)
		}
	}
	return maxProfit
}

func main() {
	values := []int{7, 6, 4, 3, 1}

	profit := bestTimeToBuyStock(values)

	fmt.Println(profit)
}
