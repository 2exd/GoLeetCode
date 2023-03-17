package greedy

/*121-买卖股票的最佳时机*/
func maxProfit(prices []int) int {
	min := prices[0]
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > res {
			res = prices[i] - min
		}
		if min > prices[i] {
			min = prices[i]
		}
	}
	return res
}
