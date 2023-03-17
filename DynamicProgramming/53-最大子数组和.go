package main

/*53.最大子数组和*/
// func maxSubArray(nums []int) int {
// 	l := len(nums) - 1
// 	res, sum := nums[1], 0
// 	for i := 0; i < l; i++ {
//
// 		for j := i; j <= l; j++ {
// 			sum += nums[j]
// 			res = findMax53(sum, res)
// 		}
//
// 	}
// 	return res
// }

func findMax53(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxSubArray(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	dp[0] = nums[0]
	res := nums[0]
	for i := 1; i < l; i++ {
		dp[i] = findMax53(dp[i-1]+nums[i], nums[i])
		res = findMax53(res, dp[i])
	}
	return res
}
