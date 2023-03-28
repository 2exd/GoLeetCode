package main

import (
	"strings"
)

/*
题干：
	给定一个整数x， 请你构造一个仅由 ‘r’、'e'、'd'三种字符组成的 字符串，其中回文子串的数量恰好为x。
	字符串的长度不得超过10的5次方。

输入描述
	一个正整数x。
	1<x<10的9次方

输出描述
	输出一个仅由'r'、'e'、'd'三种字符组成的字符串，长度不得超过10的五次方。有多解时输出任意即可。
*/
// func main() {
// 	var x int
// 	fmt.Scan(&x)
// 	if x == 1 {
// 		fmt.Println("r")
// 		return
// 	}
// 	// 解法1 暴力，答案覆盖不全
// 	fmt.Println(redPalindrome(x))
// }

func redPalindrome(x int) string {
	palindromeCount := (x + 1) / 2
	return strings.Repeat("red", palindromeCount)[:x]
}
