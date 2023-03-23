package main

import "fmt"

/*
春节期间小明使用微信收到很多个红包，非常开心。在查看领取红包记录时发现，某个红包金额出现的次数超过了红包总数的一半。请帮小明找到该红包金额。写出具体算法思路和代码实现，要求算法尽可能高效。
给定一个红包的金额数组 gifts 及它的大小 n ，请返回所求红包的金额。
若没有金额超过总数的一半，返回0。
*/
func gift() int {
	var gifts []int
	var n, num int

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		num := num
		_, err := fmt.Scan(&num)
		if err != nil {
			return 0
		} else {
			gifts = append(gifts, num)

		}
	}

	m := make(map[int]int, n)
	for i := 0; i < n; i++ {
		m[gifts[i]]++
		if m[gifts[i]] > n/2 {
			return gifts[i]
		}
	}
	return 0
}

// func main() {
// 	fmt.Println(gift())
// }
