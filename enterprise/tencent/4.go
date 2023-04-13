package main

var (
	path4  []int
	res4   int
	input4 int
)

// func main() {
// 	fmt.Scan(&input4)
// 	nums := make([]int, input4)
// 	for i := 0; i < input4; i++ {
// 		fmt.Scan(&nums[i])
// 	}
// 	// fmt.Print("4")
// 	fmt.Printf("%d\n", solution4(nums))
// }

func solution4(nums []int) int {
	if input4 == 0 {
		return 0
	}
	path4 = make([]int, input4)
	dfs4(nums, 0)
	return res4
}

func dfs4(nums []int, start int) {
	if start >= input4 {
		return
	}

	if computeMul(path4) == computeXor(path4) {
		res4++
	}
	for i := start; i < input4; i++ {
		path4 = append(path4, nums[i])
		dfs4(nums, start+1)
		path4 = path4[:len(path4)-1]
	}
}

func computeMul(nums []int) int {
	mul := 1
	for _, num := range nums {
		mul *= num
	}
	return mul
}

func computeXor(nums []int) int {
	xor := nums[0]
	if len(nums) <= 1 {
		return xor
	}
	for i := 1; i < len(nums); i++ {
		xor ^= nums[i]
	}
	return xor
}
