package main

import (
	"sort"
)

func coinChange(coins []int, amount int) int {
	res322 := 0

	if amount == 0 {
		return 0
	}

	sort.Ints(coins)

	for i := len(coins) - 1; i >= 0; i-- {
		for amount >= coins[i] {
			amount -= coins[i]
			res322++
		}
	}

	if amount > 0 {
		return -1
	}
	return res322
}

// func main() {
// 	fmt.Print(coinChange([]int{1, 2, 5}, 11))
// }
