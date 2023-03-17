package main

var (
	path46 []int
	res46  [][]int

	state46 []bool
)

func permute(nums []int) [][]int {
	path46 = make([]int, 0, len(nums))
	state46 = make([]bool, len(nums))
	res46 = make([][]int, 0, len(nums))
	dfs46(nums)
	return res46
}

func dfs46(nums []int) {
	if len(nums) == len(path46) {
		tmp := make([]int, len(nums))
		copy(tmp, path46)
		res46 = append(res46, tmp)
	}

	for i := 0; i < len(nums); i++ {
		if !state46[i] {
			path46 = append(path46, nums[i])
			state46[i] = true
			dfs46(nums)
			state46[i] = false
			path46 = path46[:len(path46)-1]
		}
	}
}

// func main() {
// 	permute([]int{1, 2, 3})
// }
