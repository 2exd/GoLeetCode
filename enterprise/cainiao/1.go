package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
题目描述：
	我们定义一棵树是好树当且仅当它的红色节点数量大于蓝色节点数量。
	现在给定一棵好树，你需要删除一条边，使得形成的两个子树均为好树，请你求出有多少种删边方法

输入：
	3
	RBR
	1 2
	2 3

输出：
	2

思路:
	dfs 遍历，每次计算以当前节点为根的树的红蓝节点个数。
	并判断此情况下的两棵树是否同时为好树。

*/
type node struct {
	color byte
	edges []int
}

var (
	numRed  int // 所有红色节点个数
	numBlue int // 所有蓝色节点个数
)

func dfs(u int, parent int, nodes []node, red, blue []int, ans *int) {
	// 先将其本身的颜色统计一下
	if nodes[u].color == 'R' {
		red[u]++
	} else {
		blue[u]++
	}
	// 遍历邻接表
	for _, v := range nodes[u].edges {
		// dfs
		if v != parent {
			dfs(v, u, nodes, red, blue, ans)
			// 统计子树中红蓝节点个数
			red[u] += red[v]
			blue[u] += blue[v]
		}
	}
	// 判断两棵树是否都为好树
	if red[u]-blue[u] > 0 && numRed-red[u] > numBlue-blue[u] {
		*ans += 1
	}
}

func main() {
	var n int
	fmt.Scanln(&n)

	nodes := make([]node, n)
	colors := make([]byte, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%c", &colors[i])
		if colors[i] == 'R' {
			numRed++
		} else {
			numBlue++
		}
		nodes[i].color = colors[i]
	}
	scanner := bufio.NewScanner(os.Stdin)
	// 返回每个以空格分隔的文本单词，并删除周围的空格
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n-1; i++ {
		var u, v int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &u)
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &v)
		u--
		v--
		nodes[u].edges = append(nodes[u].edges, v)
		nodes[v].edges = append(nodes[v].edges, u)
	}

	red := make([]int, n)
	blue := make([]int, n)
	ans := 0
	dfs(0, -1, nodes, red, blue, &ans)
	fmt.Println(ans)
}
