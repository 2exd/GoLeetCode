package main

/*
518-零钱兑换 II
*/
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {

		}
	}
	return dp[amount]
}

