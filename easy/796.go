package main

import (
	"strings"
)

/*796. 旋转字符串*/

func rotateString(s, goal string) bool {
	n := len(s)
	if n != len(goal) {
		return false
	}
next:
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// 判断是否符合
			if s[(i+j)%n] != goal[j] {
				// 不符合就直接进入下一次循环
				continue next
			}
		}
		return true
	}
	return false
}

/*
复杂度分析
时间复杂度：O(n^2)，其中 n 是字符串 s 的长度。我们需要双重循环来判断。
空间复杂度：O(1)。仅使用常数空间。
。
*/

/*方法二：搜索子字符串*/
/*
首先，如果 s 和 goal 的长度不一样，那么无论怎么旋转，s 都不能得到 goal，返回 \false。
字符串 s + s 包含了所有 s 可以通过旋转操作得到的字符串，只需要检查 goal 是否为 s + s 的子字符串即可
*/
func goodrotateString(s, goal string) bool {
	// return len(s) == len(goal) && strings.Contains(s+s, goal)
	b := string(append([]byte(s), s...))
	return len(s) == len(goal) && strings.Contains(b, goal)
}

/*复杂度分析
时间复杂度：O(n)，其中 n 是字符串 s 的长度。KMP 算法搜索子字符串的时间复杂度为 O(n)，其他搜索子字符串的方法会略有差异。
空间复杂度：O(n)，其中 n 是字符串 s 的长度。KMP 算法搜索子字符串的空间复杂度为 O(n)，其他搜索子字符串的方法会略有差异。
。*/

func kmprotateString(s, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if strStrV2(s+s, goal) == -1 {
		return false
	}
	return true
}

// func main()  {
//	var s = "abcde"
//	var goal = "cdeab"
//	fmt.Println(rotateString(s,goal))
// }

// 生成 next 数组
func generateNext(p string) []int {
	m := len(p)
	next := make([]int, m, m)
	next[0] = -1
	next[1] = 0
	i, j := 0, 1 // 前缀子串、后缀子串起始位置
	// 因为是通过最长可匹配前缀子串计算，所以 j 的值最大不会超过 m-1
	for j < m-1 {
		if i == -1 || p[i] == p[j] {
			i++
			j++
			// 设置当前最长可匹配前缀子串结尾字符下标
			next[j] = i
		} else {
			// 如果 p[i] != p[j]，回到上一个最长可匹配前缀子串结尾字符下标
			i = next[i]
		}
	}
	return next
}

// KMP 算法实现函数
func kmpSearch(s, p string) int {
	n := len(s)             // 主串长度
	m := len(p)             // 模式串长度
	next := generateNext(p) // 生成 next 数组
	i, j := 0, 0
	for i < n && j < m {
		// 如果主串字符和模式串字符不相等，
		// 更新模式串坏字符下标位置为好前缀最长可匹配前缀子串尾字符下标
		// 然后从这个位置重新开始与主串匹配
		// 相当于前面提到的把模式串向后移动 j - k 位
		if j == -1 || s[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == m {
		// 完全匹配，返回下标位置
		return i - j
	} else {
		return -1
	}
}

// 基于 KMP 算法实现查找字符串子串函数
func strStrV2(haystack, needle string) int {
	// 子串长度=0
	if len(needle) == 0 {
		return 0
	}
	// 主串长度=0，或者主串长度小于子串长度
	if len(haystack) == 0 || len(haystack) < len(needle) {
		return -1
	}
	// 子串长度=1时单独判断
	if len(needle) == 1 {
		i := 0
		for ; i < len(haystack); i++ {
			if haystack[i] == needle[0] {
				return i
			}
		}
		return -1
	}

	// 其他情况走 KMP 算法
	return kmpSearch(haystack, needle)
}
