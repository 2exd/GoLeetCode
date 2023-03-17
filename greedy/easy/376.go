package main

/*376. 摆动序列*/

// 贪心算法
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	ans := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		ans = 2
	}
	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
			ans++
			prevDiff = diff
		}
	}
	return ans
}

// 动态规划
func wiggleMaxLengthDynamic(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	dp := make([][2]int, n)
	//  i 0 作为波峰的最大长度
	//  i 1 作为波谷的最大长度
	dp[0][0] = 1
	dp[0][1] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] { //nums[i]为波谷
				dp[i][1] = max(dp[i][1], dp[j][0]+1)
			}
			if nums[j] < nums[i] { //nums[i]为波峰 或者相等
				dp[i][0] = max(dp[i][0], dp[j][1]+1)
			}
			if nums[j] == nums[i] { //添加一种情况，nums[i]为相等
				dp[i][0] = max(dp[i][0], dp[j][0]) //波峰
				dp[i][1] = max(dp[i][1], dp[j][1]) //波谷
			}
		}
	}
	return max(dp[n-1][0], dp[n-1][1])
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
