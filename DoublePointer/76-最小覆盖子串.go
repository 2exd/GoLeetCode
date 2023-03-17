package main

import (
	"math"
)

/*76. 最小覆盖子串*/
func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}

	// 初始化哈希表
	for i := 0; i < len(t); i++ {
		// record elements number in string t
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	// 检查是否包含t的所有字符
	check := func() bool {
		for k, v := range ori {
			// element number smaller , false
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				// 更新最小子串长度
				len = r - l + 1
				// 更新左右指针位置
				ansL, ansR = l, l+len
			}
			// pop, till do not contain string t
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

// func main() {
// 	fmt.Print(runtime.NumCPU())
// }
