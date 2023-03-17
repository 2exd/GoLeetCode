package main

func minCostClimbingStairs(cost []int) int {
	f := make([]int, len(cost)+1)
	f[0], f[1] = 0, 0
	for i := 2; i <= len(cost); i++ {
		f[i] = min746(f[i-1]+cost[i-1], f[i-2]+cost[i-2])
	}
	return f[len(cost)]
}
func min746(a, b int) int {
	if a < b {
		return a
	}
	return b
}
