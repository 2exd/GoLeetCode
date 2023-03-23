package main

import (
	"math"
)

/*
权值 = 字符种数 * 字符串长度
返回子字符串的最大权值，要求尽可能小
*/
var (
	path   []string
	res    []string
	resMax int
)

func restoreIpAddresses(s string, k int) int {
	path, res = make([]string, 0, len(s)), make([]string, 0)
	resMax = math.MaxInt
	dfs(s, 0, k)
	return resMax
}
func dfs(s string, start, k int) {
	if len(path) >= k { // 够k段后就不再继续往下递归
		if start == len(s) && len(path) == k {
			res = append(res, s)
			// 计算权值
			weight := computeMaxWeight(path)
			// 更新全局最大值
			if weight < resMax {
				resMax = weight
			}
		}
		return
	}
	for i := start; i < len(s); i++ {
		str := s[start : i+1]
		path = append(path, str) // 符合条件的就进入下一层
		dfs(s, i+1, k)
		path = path[:len(path)-1]
	}
}
func computeMaxWeight(substrs []string) int {
	weightMax := 0

	for _, substr := range substrs {
		m := make(map[int32]int)
		weight := 0
		for _, v := range substr {
			m[v]++
		}
		weight = len(m) * len(substr)
		if weight > weightMax {
			weightMax = weight
		}
	}
	return weightMax
}

// func main() {
// 	fmt.Print(restoreIpAddresses("ababbbb", 2))
// }
