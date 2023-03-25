package main

/*821. 字符的最短距离*/
func shortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)

	// 正向遍历
	idx := -n
	for i, ch := range s {
		if byte(ch) == c {
			// c的坐标
			idx = i
		}
		// 第一次遍历更新距离
		ans[i] = i - idx
	}

	// 反向遍历
	idx = n * 2
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			idx = i
		}
		ans[i] = min(ans[i], idx-i)
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
