package main

/*718-最长重复子数组*/
func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	res := 0
	dp := make([]int, n+1)
	for i := 1; i <= m; i++ {
		for j := n; j >= 1; j-- {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = 0 // 注意这里不相等要赋值为0，供下一层使用
			}
			if dp[j] > res {
				res = dp[j]
			}
		}
	}
	return res
}
