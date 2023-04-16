package main

import (
	"fmt"
)

/*
	数组中取两个元素，返回使得乘积最大的组合数（有序）
*/
func main() {
	n := 0
	fmt.Scan(&n)
	nums := make([]int, n)
	m := make(map[int]int)
	max1, max2 := 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
		m[nums[i]]++
		if nums[i] > max1 {
			max2 = max1
			max1 = nums[i]
		}
	}
	ans := m[max1]
	ans2 := m[max2]
	if ans >= 2 {
		fmt.Print(ans * (ans - 1))
	} else if ans2 >= 2 {
		fmt.Print(ans2 * 2)
	} else {
		fmt.Print(2)
	}
}
