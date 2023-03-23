package main

import "fmt"

/*
从左下到右上斜着打印二维数组
*/
func printMatrix(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])

	for i := n - 1; i >= 0; i-- {
		k := i
		for j := 0; j < m; j++ {
			if k >= n {
				break
			}
			fmt.Printf("%d,", matrix[k][j])
			k++
		}
	}

	for j := 1; j < m; j++ {
		k := j
		for i := 0; i < n; i++ {
			if k >= m {
				break
			}
			fmt.Printf("%d,", matrix[i][k])
			k++
		}
	}

}

// func main() {
// 	// matrix := [][]int{
// 	// 	{1, 2, 3},
// 	// 	{4, 5, 6},
// 	// 	{7, 8, 9},
// 	// }
//
// 	// matrix := [][]int{
// 	// 	{1, 2},
// 	// 	{4, 5},
// 	// 	{7, 8},
// 	// }
//
// 	matrix := [][]int{
// 		{1, 2, 3}}
//
// 	// matrix := [][]int{
// 	// 	{1, 2, 3, 4, 5},
// 	// 	{6, 7, 8, 9, 10},
// 	// 	{11, 12, 13, 14, 15},
// 	// 	{16, 17, 18, 19, 20},
// 	// 	{21, 22, 23, 24, 25},
// 	// }
//
// 	printMatrix(matrix)
// }
