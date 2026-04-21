package main

import "fmt"

func maxProfit(prices []int) int {

	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		purchasedPrice := prices[i]
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - purchasedPrice
			if maxProfit < profit {
				maxProfit = profit
			}

		}
	}
	return maxProfit

}
func main() {
	fmt.Println(maxProfit([]int{10, 1, 5, 6, 7, 1}))
}
