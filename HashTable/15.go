package main

import (
	"sort"
)

/*15. 三数之和 medium*/

/*
双指针法
时间复杂度：O(n^2)。
*/
func threeSum(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	res := [][]int{}
	length := len(nums)
	// 全是正数
	if nums[0] > 0 {
		return res
	}

	for i := 0; i < length-2; i++ {
		n1 := nums[i]
		// equal，next loop
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, length-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if nums[l]+nums[r]+n1 == 0 {
				res = append(res, []int{n1, n2, n3})
				// 直到新的左指针值不等于n2
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

// func main() {
// 	fmt.Print(threeSum([]int{0, 0, 0}))
// }
