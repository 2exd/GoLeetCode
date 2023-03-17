package main

import "fmt"

/*54-螺旋矩阵*/
func spiralOrder(matrix [][]int) []int {
	n := len(matrix[0])
	m := len(matrix)
	top, bottom := 0, m-1
	left, right := 0, n-1
	count := m * n
	var res []int

	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
			count--
		}
		if count == 0 {
			return res
		}
		top++
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
			count--
		}
		if count == 0 {
			return res
		}
		right--
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
			count--
		}
		if count == 0 {
			return res
		}
		bottom--
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
			count--
		}
		if count == 0 {
			return res
		}
		left++
	}
	// return res
}

func main() {
	fmt.Print(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}
