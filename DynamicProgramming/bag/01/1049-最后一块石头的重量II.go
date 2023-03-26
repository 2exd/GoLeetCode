package main

/*1049-最后一块石头的重量II*/
func lastStoneWeightII(stones []int) int {
	// 15001 = 30 * 1000 /2 +1
	dp := make([]int, 15001)
	// 求target
	sum := 0
	for _, v := range stones {
		sum += v
	}
	// 背包容量为 target 情况下，选取的石头的最大重量和
	target := sum / 2
	// 遍历顺序
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			// 推导公式
			dp[j] = max1049(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	// sum/2 向下取整 => sum-dp[target] 一定是大于等于 dp[target] 的
	// (sum - dp[target]) - dp[target] => sum - 2*dp[target]
	return sum - 2*dp[target]
}

func max1049(a, b int) int {
	if a > b {
		return a
	}
	return b
}
