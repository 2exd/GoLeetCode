package main

var (
	res216  [][]int
	path216 []int
)

func combinationSum3(k int, n int) [][]int {
	res216, path216 = make([][]int, 0), make([]int, 0, k)
	dfs216(k, n, 1, 0)
	return res216
}

func dfs216(k, n int, start int, sum int) {
	if len(path216) == k {
		if sum == n {
			tmp := make([]int, k)
			copy(tmp, path216)
			res216 = append(res216, tmp)
		}
		return
	}
	for i := start; i <= 9; i++ {
		if sum+i > n || 9-i+1 < k-len(path216) {
			break
		}
		path216 = append(path216, i)
		dfs216(k, n, i+1, sum+i)
		path216 = path216[:len(path216)-1]
	}
}
