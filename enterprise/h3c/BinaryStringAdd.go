package main

import (
	"strconv"
)

/*
67. 二进制求和
*/

func binaryAdd(a, b string) string {
	carry := 0
	ans := ""
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)
	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
