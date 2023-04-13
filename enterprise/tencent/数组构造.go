package main

import "fmt"

/*
题目描述：
	给定正整数n,返回一个大小为 n*(n+1)/2 的数组，其中任意两个相邻元素都不相等。

输入：
	3
输出：
	3,2,3,2，3,1
*/

// func main() {
// 	var n int
// 	fmt.Scanln(&n)
// 	constructArray(n)
// }
func constructArray(n int) {
	var arr []int

	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			arr = append(arr, n-j)
		}
	}
	fmt.Print(arr)
}
