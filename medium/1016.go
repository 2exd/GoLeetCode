package main

import (
	"strings"
)

/*1016. 子串能表示从 1 到 N 数字的二进制串*/
// func getBi(a int) (res string) {
//	if a == 0 {
//		return "0"
//	}
//	for a > 0 {
//		if a&1 == 1 {
//			res = "1" + res
//		} else {
//			res = "0" + res
//		}
//		a >>= 1
//	}
//	return res
// }

// func getBi(a int)  string {
//	var builder strings.Builder
//	if a == 0 {
//		return "0"
//	}
//	for a > 0 {
//		if a&1 == 1 {
//			//res = "1" + res
//			builder.WriteString("1")
//		} else {
//			//res = "0" + res
//			builder.WriteString("0")
//		}
//		a >>= 1
//	}
//	return builder.String()
// }

func getBi(a int) string {
	buf := make([]byte, 0)
	if a == 0 {
		return "0"
	}
	for a > 0 {
		if a&1 == 1 {
			// res = "1" + res
			buf = append([]byte{'1'}, buf...)
		} else {
			// res = "0" + res
			buf = append([]byte{'0'}, buf...)
		}
		a >>= 1
	}
	return string(buf)
}

func queryString(s string, n int) bool {
	limit := (1 << 9) - 1
	if n > limit {
		return false
	}
	for i := 1; i <= n; i++ {
		if !strings.Contains(s, getBi(i)) {
			return false
		}
	}
	return true
}

/*
题目的入参范围看起来很大，n可以达到1e9，即使用高效的字符串匹配O(mn)下时间复杂度也达到了1e12。
但是实际上我们可以发现由于待匹配的子串是连续的，即使存在一些子串能相互复用一部分，粗略估计一下
范围显然当n超过2^9-1即二进制111111111时s1000的长度就已经不可能容纳所有的字串，时间复杂度仅为5e6,
直接用strings包的contains判断击败100%
*/
