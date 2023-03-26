package main

import "fmt"

func findMaxForm(strs []string, m int, n int) int {
	// 定义数组
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	// 遍历
	for i := 0; i < len(strs); i++ {
		zeroNum, oneNum := 0, 0
		// 计算0,1 个数
		// 或者直接strings.Count(strs[i],"0")
		for _, v := range strs[i] {
			if v == '0' {
				zeroNum++
			}
		}
		oneNum = len(strs[i]) - zeroNum
		// 从后往前 遍历背包容量
		for j := m; j >= zeroNum; j-- {
			for k := n; k >= oneNum; k-- {
				// 推导公式
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

func main() {

	for i := range [100]struct{}{} {
		fmt.Println(i)
	}
	// findMaxForm([]string{"10", "0", "1"}, 1, 1)
	// findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3)
	// findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 3, 4)
}
