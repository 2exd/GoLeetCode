package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
题目描述
	小红拿到了棵树，每个节点被染成了红色或者蓝色。
	小红定义每条边的权值为:删除这条边时，形成的两个子树的同色连通块数量之差的绝对值。
	小红想知道，所有边的权值之和是多少?

输入描述
	第一行输入一个正整数n,代表节点的数量。
	第二行输入一个长度为n且仅由'R"和B，两种字符组成的字符串。
	第i个字符为'R"代表i号节点被染成红色，为'B'则被染成蓝色。
	接下来的n-1行，每行输入两个正整数u和U，代表节点么和节 点U有一条边相连
	1≤n≤200000

输出描述
	一个正整数，代表所有节点的权值之和。
*/
type node struct {
	color byte
	edges []int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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
	// 计算权值
	*ans += abs(red[u] - blue[u])
}

func main() {
	var n int
	fmt.Scanln(&n)

	nodes := make([]node, n)
	colors := make([]byte, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%c", &colors[i])
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
