package main

/*122. 买卖股票的最佳时机 II*/

// 贪心算法
func maxProfit(prices []int) int {
	var sum int
	for i := 1; i < len(prices); i++ {
		// 累加每次大于0的交易
		if prices[i]-prices[i-1] > 0 {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}
