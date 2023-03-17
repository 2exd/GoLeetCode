package dfs

/*386. 字典序排数*/
func lexicalOrder(n int) []int {
	ans := make([]int, n)
	num := 1
	for i := range ans {
		ans[i] = num
		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++
		}
	}
	return ans
}

/*
复杂度分析
时间复杂度：O(n)，其中 n 为整数的数目。获取下一个字典序整数的最坏时间复杂度为 O(logn)，但 while 循环的迭代次数与 number 的末尾连续的 9 的数目有关，
在整数区间 [1,n] 中，末尾连续的 9 的数目为 k 的整数不超过 ，总时间复杂度为 O(n)。
空间复杂度：O(1)。返回值不计入空间复杂度。

*/
