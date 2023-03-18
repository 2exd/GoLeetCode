package main

import "fmt"

/*
小美在玩一项游戏，该游戏的目标是尽可能抓获敌人。
敌人的位置将被一个二维坐标 (x,y) 所描述。
小美有一个全屏技能，该技能能一次性将若干故人一次性捕获。
捕获的敌人之间的横坐标的最大差值不能大于A，纵坐标的最大差值不能大于B
现在给出所有敌人的坐标，你的任务是计算小美一次任最多能使用技能甫获多少敌人。
输大描述：
第一行三个整数N,A，B
表示共有N个敌人，小美的全屏技能的参数A和参数B
接下来N行，每行两个数字x，y，描达一个敌人所在的坐标。
1<=N<=500   1<=A,B<=1000. 1<=x,y<=1000
输出描述
一行，一个整数表示小美使用技能单次所可以捕获的最多数量。
*/

// func main() {
// 	buhuo()
// }

func buhuo() int {
	// N := 1005
	var n, a, b int
	fmt.Scanf("%d %d %d\n", &n, &a, &b)
	g := make([][]int, n+1)
	s := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		g[i] = make([]int, n+1, n+1)
		s[i] = make([]int, n+1, n+1)
	}
	// 初始化图
	for i := 0; i < n; i++ {
		x, y := 0, 0
		fmt.Scanf("%d %d\n", &x, &y)
		g[x][y]++
	}

	// 初始化dp数组 代表[0~i][0~j]区域的敌人总数量
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			s[i][j] = s[i-1][j] + s[i][j-1] - s[i-1][j-1] + g[i][j]
		}
	}

	res := -1
	for i := 1; i+a < n; i++ {
		for j := i; j+b < n; j++ {
			u := i + a
			v := j + b
			// [i~u][j~v]区域能捕获的敌人数量
			res = findMAxBuoy(res, s[u][v]+s[i-1][j-1]-s[i-1][v]-s[u][j-1])
		}
	}
	fmt.Println(res)
	return res
}

func findMAxBuoy(x, y int) int {
	if x > y {
		return x
	}
	return y
}
