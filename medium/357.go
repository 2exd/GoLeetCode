package main

/*357. 统计各位数字都不同的数字个数*/
/*
0 <= n <= 8
首位可选数字为9个
第二位开始可选数字从9递减
*/
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	ans, cur := 10, 9
	for i := 0; i < n-1; i++ {
		cur *= 9 - i
		ans += cur
	}
	return ans
}
