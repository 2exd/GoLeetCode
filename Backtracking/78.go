package main

var (
	path78 []int
	res78  [][]int
)

func subsets(nums []int) [][]int {
	res78, path78 = make([][]int, 0), make([]int, 0, len(nums))
	dfs78(nums, 0)
	return res78
}
func dfs78(nums []int, start int) {
	tmp := make([]int, len(path78))
	copy(tmp, path78)
	res78 = append(res78, tmp)

	for i := start; i < len(nums); i++ {
		path78 = append(path78, nums[i])
		dfs78(nums, i+1)
		path78 = path78[:len(path78)-1]
	}
}
