package main

/*剑指 Offer II 107. 矩阵中的距离*/
func myupdateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	// 四个方向
	dx := []int{1, 0, -1, 0}
	dy := []int{0, -1, 0, 1}
	var queue [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				// 0元素先加入到数组中
				queue = append(queue, []int{i, j})
			}
		}
	}
	for len(queue) > 0 {
		x := queue[0][0]
		y := queue[0][1]
		queue = queue[1:]
		for k := 0; k < 4; k++ {
			nx := x + dx[k]
			ny := y + dy[k]
			if nx >= 0 && ny >= 0 && nx < m && ny < n && mat[nx][ny] == 1 {
				if res[nx][ny] == 0 || res[x][y]+1 < res[nx][ny] {
					res[nx][ny] = res[x][y] + 1
					queue = append(queue, []int{nx, ny})
				}
			}
		}
	}
	return res
}

/*高分做法*/
func updateMatrix(mat [][]int) [][]int {
	// 初始化
	m, n := len(mat), len(mat[0])
	ans := make([][]int, m)
	var q [][2]int
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				ans[i][j] = 0
				q = append(q, [2]int{i, j})
			} else {
				ans[i][j] = -1
			}
		}
	}
	// 四个方向
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]
		for _, d := range dirs {
			r, c := x+d[0], y+d[1]
			if r >= 0 && r < m && c >= 0 && c < n && ans[r][c] == -1 {
				ans[r][c] = ans[x][y] + 1
				q = append(q, [2]int{r, c})
			}
		}
	}
	return ans
}

//func main() {
//	//arr := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
//	//arr := [][]int{{0,0,0},{0,1,0},{1,1,1}}
//	arr := [][]int{{0}, {1}}
//	fmt.Println(updateMatrix(arr))
//}
