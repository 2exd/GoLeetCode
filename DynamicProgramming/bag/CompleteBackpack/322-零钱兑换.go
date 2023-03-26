package main

import (
	"math"
)

/*322-零钱兑换*/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] && dp[i-coins[j]] != math.MaxInt32 {
				dp[i] = findMin322(dp[i-coins[j]]+1, dp[i])
			}
		}
	}
	// 没找到能装满背包的, 就返回-1
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func findMin322(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//
// func main() {
// 	coinChange([]int{1, 2, 5}, 11)
// }
