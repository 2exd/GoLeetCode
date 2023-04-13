package main

import "fmt"

/*
题目描述：
	给定一个正整数，你需要删除其中连续的一段数字，使得它变成15的倍数，请你求出有多少种删除方式
	ps：删除的位置不通，即可标记为不同的删除方式，可以存在前导0

输入：
	12120

输出：
	4

思路：
	如何判断为15的倍数?
		先判断是否能被5整除，即个位数为0或5
		再判断是否能被3整除，即所有位数加起来为3的整数倍
	因为是连续删除，可以采用双指针的做法，双指针包含的区域为删除的位置。
*/
func main() {
	var num int
	fmt.Scanln(&num)
	solution3(num)
}

func solution3(num int) {
	res := 0
	tmp := num
	sum := 0
	var nums []int
	for tmp != 0 {
		nums = append(nums, tmp%10)
		sum += tmp % 10
		tmp /= 10
	}
	if sum%3 == 0 && (nums[0] == 0 || nums[0] == 5) {
		res++
	}

	// 如何判断尾数位置？
	// 如果 l == 0， 那么 r+1 为尾数的位置
	// 如果 l != 0, 那么 0 为尾数的位置
	for l := 0; l < len(nums); l++ {
		tmp = sum
		for r := l; r < len(nums); r++ {
			tmp -= nums[r]
			if l != 0 {
				if nums[0] != 0 && nums[0] != 5 {
					break
				}
			} else {
				if r+1 < len(nums) {
					if nums[r+1] != 0 && nums[r+1] != 5 {
						continue
					}
				}
			}
			if tmp%3 == 0 && tmp != 0 {
				res++
			}
		}
	}
	fmt.Print(res)
}
