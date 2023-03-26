package main

import "fmt"

var (
	path5     []int
	count5    int
	input5, k int
)

func main() {
	fmt.Scan(&input5, &k)
	nums := make([]int, input5)
	for i := 0; i < input5; i++ {
		fmt.Scan(&nums[i])
	}
	// fmt.Print("4")
	solution5(nums)
	fmt.Print(count5)
}
func solution5(nums []int) {
	if input5 == 0 {
		return
	}
	path5 = make([]int, 0, input5)
	dfs5(nums, 0, k)
}
func dfs5(nums []int, start int, k int) {
	if start == len(nums) {
		return
	}
	for i := start; i < len(nums); i++ {
		path5 = append(path5, nums[i])
		if judge(nums, k) && len(path5) != len(nums) {
			count5++
		}
		dfs5(nums, i+1, k)
		path5 = path5[:len(path5)-1]
	}
}

func judge(nums []int, k int) bool {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		m = gcd(m, nums[i])
	}
	return m == k
}

func gcd(m, n int) int {
	if m < n {
		tmp := m
		m = n
		n = tmp
	}
	if n == 0 {
		return m
	}

	return gcd(n, m%n)
}
