package main

import (
	"fmt"
	"sort"
)

type Plan struct {
	x, y int // 基站的编号
	z    int // 架设光纤的成本
}

var fa []int

func find(x int) int {
	if fa[x] == x {
		return x
	}
	fa[x] = find(fa[x])
	return fa[x]
}

func merge(x, y int) {
	fx := find(x)
	fy := find(y)
	if fx != fy {
		fa[fx] = fy
	}
}

func station() {
	var N, M int
	fmt.Scan(&N, &M)

	// 初始化并查集
	fa = make([]int, N+1)
	for i := 1; i <= N; i++ {
		fa[i] = i
	}

	var plans []Plan

	for i := 0; i < M; i++ {
		var x, y, z, p int
		fmt.Scan(&x, &y, &z, &p)
		if p == 1 {
			merge(x, y)
		} else {
			plans = append(plans, Plan{x, y, z})
		}
	}

	// 按成本对计划排序
	sort.Slice(plans, func(i, j int) bool {
		return plans[i].z < plans[j].z
	})

	// Kruskal算法处理
	totalCost := 0
	for _, plan := range plans {
		if find(plan.x) != find(plan.y) {
			merge(plan.x, plan.y)
			totalCost += plan.z
		}
	}

	// 检查所有基站是否联通
	root := find(1)
	allConnected := true
	for i := 1; i <= N; i++ {
		if find(i) != root {
			allConnected = false
			break
		}
	}

	if allConnected {
		fmt.Println(totalCost)
	} else {
		fmt.Println(-1)
	}
}

// func main() {
// 	station()
// }
