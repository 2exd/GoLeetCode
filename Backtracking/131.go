package main

var (
	path131 []string // 放已经回文的子串
	res131  [][]string
)

func partition(s string) [][]string {
	path131, res131 = make([]string, 0), make([][]string, 0)
	dfs131(s, 0)
	return res131
}

func dfs131(s string, start int) {
	if start == len(s) { // 如果起始位置等于s的大小，说明已经找到了一组分割方案了
		tmp := make([]string, len(path131))
		copy(tmp, path131)
		res131 = append(res131, tmp)
		return
	}
	for i := start; i < len(s); i++ {
		str := s[start : i+1]
		if isPalindrome(str) { // 是回文子串
			path131 = append(path131, str)
			dfs131(s, i+1)                     // 寻找i+1为起始位置的子串
			path131 = path131[:len(path131)-1] // 回溯过程，弹出本次已经填在的子串
		}
	}
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
