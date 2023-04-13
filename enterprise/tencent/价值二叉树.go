package main

import (
	"fmt"
)

/*
题干：
	计算二叉树每个节点的权值
	权值=子树节点乘积末尾0的个数
*/

func main() {
	valueBinaryTree([]int{2, 5, 10, 0, 0, 4, 5})
}

func valueBinaryTree(nums []int) {
	length := len(nums)
	ans := make([]int, length)
	for i := (length / 2) - 1; i >= 0; i-- {
		if nums[2*i+1] != 0 && nums[2*i+2] != 0 {
			nums[i] = nums[2*i+1] * nums[2*i+2] * nums[i]
		}
		ans[i] = countZero(nums[i])
	}
	fmt.Print(ans)
}

func countZero(num int) int {
	count := 0
	for num >= 10 {
		num /= 10
		count++
	}
	return count
}
