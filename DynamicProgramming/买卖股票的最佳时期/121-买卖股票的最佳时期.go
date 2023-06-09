package main

/*121-买卖股票的最佳时期*/
func maxProfit121(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}
	// initialize
	dp[0][0] = -prices[0]
	dp[0][1] = 0

	for i := 1; i < length; i++ {
		dp[i][0] = max121(dp[i-1][0], -prices[i])
		dp[i][1] = max121(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[length-1][1]
}

func max121(a, b int) int {
	if a > b {
		return a
	}
	return b
}
