package main

import "math"

/*279-完全平方数*/
// func main() {
//
// }

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i, _ := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	m := int(math.Sqrt(float64(n)))
	for i := 1; i <= m; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = findMin279(dp[j-i*i]+1, dp[j])
		}
	}
	return dp[n]
}

func findMin279(x, y int) int {
	if x < y {
		return x
	}
	return y
}
