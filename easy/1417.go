package main

import "unicode"

/*819. 最常见的单词*/

func reformat(s string) string {
	sumDigit := 0
	for _, c := range s {
		if unicode.IsDigit(c) {
			sumDigit++
		}
	}

	// 如果差值大于一，不符合题目条件
	sumAlpha := len(s) - sumDigit
	if abs(sumDigit-sumAlpha) > 1 {
		return ""
	}

	flag := sumDigit > sumAlpha
	t := []byte(s)
	for i, j := 0, 1; i < len(t); i += 2 {
		if unicode.IsDigit(rune(t[i])) != flag {
			for unicode.IsDigit(rune(t[j])) != flag {
				j += 2
			}
			t[i], t[j] = t[j], t[i]
		}
	}
	return string(t)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
复杂度分析

时间复杂度：O(n)，其中 n 为字符串 s 的长度，需要遍历两遍字符串。
空间复杂度：对于字符串可变的语言为 O(1)，仅使用常量空间。而对于字符串不可变的语言需要新建一个和 s 等长的字符串，所以空间复杂度是 O(n)。

*/
