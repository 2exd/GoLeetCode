package main

func findMaxForm(strs []string, m int, n int) int {
	// 定义数组
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	// 遍历
	for _, str := range strs {
		zeroNum, oneNum := 0, 0
		// 计算 0,1 个数
		// 或者直接 strings.Count(strs[i],"0")
		for _, v := range str {
			if v == '0' {
				zeroNum++
			}
		}
		oneNum = len(str) - zeroNum
		// 从后往前遍历背包容量
		for j := m; j >= zeroNum; j-- {
			for k := n; k >= oneNum; k-- {
				// 推导公式 dp[j][k] => 0最多为j个，1最多为k个时，最大的子集的长度
				dp[j][k] = max474(dp[j][k], dp[j-zeroNum][k-oneNum]+1)
			}
		}
		// fmt.Println(dp)
	}
	return dp[m][n]
}

func max474(a, b int) int {
	if a > b {
		return a
	}
	return b
}
