package main

import "sort"

var (
	res90  [][]int
	path90 []int
)

func dfs90(start int, nums []int) {

	tmp := make([]int, len(path90))
	copy(tmp, path90)
	res90 = append(res90, tmp)

	for i := start; i < len(nums); i++ {
		// 去重
		if i != start && nums[i] == nums[i-1] {
			continue
		}
		path90 = append(path90, nums[i])
		dfs90(i+1, nums)
		path90 = path90[:len(path90)-1]
	}

}
func subsetsWithDup(nums []int) [][]int {
	res90 = make([][]int, 0, len(nums))
	path90 = make([]int, 0, len(nums))
	// sort before
	sort.Ints(nums)
	dfs90(0, nums)
	return res90
}

// func main() {
// 	subsetsWithDup([]int{1, 1, 2, 2})
// }
