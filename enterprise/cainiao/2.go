package main

import (
	"fmt"
)

/*
题目描述：
	给定一个数组，每次操作可以选择数组的一个元素，将它的二进制的某一位取反
	（假定每个元素都是32位非负整数，范围在 [0, 2^32 -1]）
	请你求出使得所有元素都相等的最小操作次数

输入：
	4
	9 9 9 9

输出：
	0

思路：
	从低位到高位，每次循环找出n个数二进制位为1的个数，记录在count中，
	如果count>= n/2，则此位需要操作 n-count 次，反之，需要操作 count 次。
	最后返回所有操作次数之和。
*/

// func main() {
// 	var n int
// 	fmt.Scanln(&n)
//
// 	nums := make([]int, n)
//
// 	for i := 0; i < n; i++ {
// 		fmt.Scanf("%d", &nums[i])
// 	}
//
// 	solution2(nums)
// }

func solution2(nums []int) {
	res := 0
	halfNum := len(nums) >> 1
	for i := 0; i < 32; i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			count += (nums[j] >> i) & 1
		}
		if count >= halfNum {
			res += len(nums) - count
		} else {
			res += count
		}
	}
	fmt.Println(res)
}
