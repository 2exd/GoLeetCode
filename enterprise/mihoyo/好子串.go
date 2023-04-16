package main

import (
	"fmt"
)

/*
	返回好子串的数目（只包含长度为1的回文子串的字符串）
*/
/*

 */

// func main() {
// 	n := 0
// 	fmt.Scan(&n)
// 	var s string
// 	fmt.Scan(&s)
// 	solution2(s)
// }

func solution2(s string) {
	sum := 0
	list := make([]string, 0)
	for i := 0; i < len(s); i++ {
		c1 := s[i]
		size := len(list)
		for j := 0; j < size; j++ {
			c2 := list[0]
			list = list[1:]
			var bytes []byte
			bytes = append(bytes, []byte(c2)...)
			bytes = append(bytes, c1)
			if isHuiWen(string(bytes)) {
				list = append(list, string(bytes))
				sum++
			}
		}
		list = append(list, string(c1))
		sum++
	}
	fmt.Print(sum)
}

// func isPalindrome(s string) bool {
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		if s[i] != s[j] {
// 			return false
// 		}
// 	}
// 	return true
// }
func isHuiWen(s string) bool {
	var temp string
	temp = temp + string(s[len(s)-1])
	// for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
	//
	// 	if s[i] != s[j] {
	// 		return false
	// 	}
	// }

	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		temp = temp + string(c)
		for j := 0; j < len(temp); j++ {
			k := len(temp) - j - 1
			if temp[k] != temp[j] {
				return true
			}
		}
	}

	return false
}

func countSubstrings(s string) int {
	n := len(s)
	t := "$#"
	for i := 0; i < n; i++ {
		t += string(s[i]) + "#"
	}
	n = len(t)
	t += "!"

	f := make([]int, n)
	iMax, rMax, ans := 0, 0, 0
	for i := 1; i < n; i++ {
		// 初始化 f[i]
		if i <= rMax {
			f[i] = min(rMax-i+1, f[2*iMax-i])
		} else {
			f[i] = 1
		}
		// 中心拓展
		for t[i+f[i]] == t[i-f[i]] {
			f[i]++
		}
		// 动态维护 iMax 和 rMax
		if i+f[i]-1 > rMax {
			iMax = i
			rMax = i + f[i] - 1
		}
		// 统计答案, 当前贡献为 (f[i] - 1) / 2 上取整
		ans += f[i] / 2
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
