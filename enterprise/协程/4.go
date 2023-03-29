package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n   int
	res int
	vis []bool
	g   [][]int
)

func findMax(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Scanln(&n)
	vis = make([]bool, n+1)
	g = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		g[i] = make([]int, n+1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	// 返回每个以空格分隔的文本单词，并删除周围的空格
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n-1; i++ {
		var u, v, w int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &u)
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &v)
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &w)

		// edge
		g[u][v] = w
	}
	dfs(0, 0)
	fmt.Println(res)
}

func dfs(cnt, sum int) {
	if cnt == 2 {
		res = findMax(sum, res)
		return
	}

	for i := 1; i <= n; i++ {
		if vis[i] {
			continue
		}
		vis[i] = true
		for j := 1; j <= n; j++ {
			if i != j && g[i][j] != 0 && !vis[j] {
				vis[j] = true
				dfs(cnt+1, sum+g[i][j])
				vis[j] = false
			}
		}
		vis[i] = false
	}
}
