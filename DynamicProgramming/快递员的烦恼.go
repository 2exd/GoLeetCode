package main

import (
	"fmt"
	"math"
)

func kuaiDiYuan() {
	var n, m int
	fmt.Scan(&n, &m)
	dist := make([][]int, n+1)
	for i := range dist {
		dist[i] = make([]int, n+1)
		for j := range dist[i] {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}

	// 客户ID到索引的映射
	idxMap := make(map[int]int)

	// 初始化投递站到每个客户的距离
	for idx := 1; idx <= n; idx++ {
		var cid, distance int
		fmt.Scan(&cid, &distance)
		dist[0][idx] = distance
		dist[idx][0] = distance
		idxMap[cid] = idx
	}

	// 初始化客户之间的距离
	for i := 0; i < m; i++ {
		var cid1, cid2, distance int
		fmt.Scan(&cid1, &cid2, &distance)
		idx1 := idxMap[cid1]
		idx2 := idxMap[cid2]
		dist[idx1][idx2] = distance
		dist[idx2][idx1] = distance
	}

	// 使用Floyd-Warshall算法计算所有点对之间的最短距离
	for k := 0; k <= n; k++ {
		for i := 0; i <= n; i++ {
			for j := 0; j <= n; j++ {
				if dist[i][k] < math.MaxInt32 && dist[k][j] < math.MaxInt32 {
					dist[i][j] = minKuaidi(dist[i][j], dist[i][k]+dist[k][j])
				}
			}
		}
	}

	// 使用动态规划解决旅行商问题
	dp := make([][]int, 1<<(n+1))
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[1][0] = 0 // 在投递站开始

	// 遍历所有状态
	for state := 0; state < (1 << (n + 1)); state++ {
		for i := 0; i <= n; i++ {
			if state>>i&1 == 1 && dp[state][i] != math.MaxInt32 {
				for last := 0; last <= n; last++ {
					dp[state|(1<<last)][last] = minKuaidi(dp[state|(1<<last)][last], dp[state][i]+dist[i][last])
				}
			}
		}
	}

	// 输出最终回到投递站的最短距离
	if dp[(1<<(n+1))-1][0] == math.MaxInt32 {
		fmt.Println(-1)
	} else {
		fmt.Println(dp[(1<<(n+1))-1][0])
	}
}

// 辅助函数：返回两个整数的最小值
func minKuaidi(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func main() {
// 	kuaiDiYuan()
// }
