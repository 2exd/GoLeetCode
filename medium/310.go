package main

import "fmt"

/*310. 最小高度树*/
/*构建图，循环遍历图，找出叶子节点。去除叶子节点。知道图中节点只剩下2个或1个。返回剩下的节点。*/

/*BFS*/

/*设 dist[x][y] 表示从节点 x 到节点 y 的距离，假设树中距离最长的两个节点为 (x,y)，它们之间的距离为 maxdist=dist[x][y]，
则可以推出以任意节点构成的树最小高度一定为minheight=⌈maxdist/2⌉，且最小高度的树根节点一定在 节点 x 到节点 y 的路径上。*/

/*且最小高度树的根节点一定存在其最长路径上。假设最长的路径的 m 个节点依次为 p1  p2 。。。 pm，最长路径的长度为 m-1，可以得到以下结论：
如果 m 为偶数，此时最小高度树的根节点为 pm/2或者pm/2 + 1 且此时最小的高度为 m/2；
如果 m 为奇数，此时最小高度树的根节点为 pm/2 + 1 且此时最小的高度为 (m-1)/2；
*/

/*
因此只需要求出路径最长的两个叶子节点即可，并求出其路径的最中间的节点即为最小高度树的根节点。可以利用以下算法找到图中距离最远的两个节点与它们之间的路径：
以任意节点 p 出现，利用广度优先搜索或者深度优先搜索找到以 p 为起点的最长路径的终点 x；
以节点 x 出发，找到以 x 为起点的最长路径的终点 y；
x 到 y 之间的路径即为图中的最长路径，找到路径的中间节点即为根节点。
*/
func bfsfindMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	// 图的邻接矩阵
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	parents := make([]int, n)
	bfs := func(start int) (x int) {
		vis := make([]bool, n)
		vis[start] = true
		q := []int{start}
		for len(q) > 0 {
			x, q = q[0], q[1:]
			for _, y := range g[x] {
				if !vis[y] {
					vis[y] = true
					parents[y] = x
					q = append(q, y)
				}
			}
		}
		return
	}
	x := bfs(0) // 找到与节点 0 最远的节点 x
	y := bfs(x) // 找到与节点 x 最远的节点 y
	fmt.Println("x ", x)
	fmt.Println("y ", y)
	fmt.Println("parents ", parents)

	var path []int
	parents[x] = -1
	for y != -1 {
		path = append(path, y)
		y = parents[y]
	}
	m := len(path)
	if m%2 == 0 {
		return []int{path[m/2-1], path[m/2]}
	}
	return []int{path[m/2]}
}

/*复杂度分析
时间复杂度：O(n)，其中 n 是为节点的个数。图中边的个数为 n-1，因此建立图的关系需要的时间复杂度为 O(n)，
通过广度优先搜索需要的时间复杂度为 O(n+n−1)，求最长路径的时间复杂度为 O(n)，因此总的时间复杂度为 O(n)。
空间复杂度：O(n)，其中 n 是节点的个数。由于题目给定的图中任何两个顶点都只有一条路径连接，因此图中边的数目刚好等于 n-1，
用邻接表构造图所需的空间刚好为 O(2×n)，存储每个节点的距离和父节点均为 O(n)，使用广度优先搜索时，队列中最多有 n 个元素，所需的空间也为 O(n)，因此空间复杂度为 O(n)。
*/

/*DFS*/
/*首先找到距离节点 0 的最远节点 x，然后找到距离节点 x 的最远节点 y，然后找到节点 x 与节点 y 的路径，然后找到根节点*/
func dfsfindMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	parents := make([]int, n)
	maxDepth, node := 0, -1
	var dfs func(int, int, int)
	dfs = func(x, pa, depth int) {
		if depth > maxDepth {
			maxDepth, node = depth, x
		}
		parents[x] = pa
		for _, y := range g[x] {
			if y != pa {
				dfs(y, x, depth+1)
			}
		}
	}
	dfs(0, -1, 1)
	maxDepth = 0
	dfs(node, -1, 1)

	path := []int{}
	for node != -1 {
		path = append(path, node)
		node = parents[node]
	}
	m := len(path)
	if m%2 == 0 {
		return []int{path[m/2-1], path[m/2]}
	}
	return []int{path[m/2]}
}

/*使用与有向图相同的拓扑排序思想，并稍加修改，可以复用于无向图排序，区别在于：
当一个节点的入度 = 1，将其推入队列；
如果邻节点是否已经访问过，则跳过；*/
func goodfindMinHeightTrees(n int, edges [][]int) []int {
	numEdges, graph, work, level := make([]int, n), make([][]int, n), make([]int, n), make([]int, n)
	for _, edge := range edges {
		// 统计数字出现次数
		numEdges[edge[0]]++
		numEdges[edge[1]]++
	}
	fmt.Println("numEdges ", numEdges)
	for i := 0; i < n; i++ {
		// 图的矩阵表示
		graph[i] = make([]int, numEdges[i])
	}
	for _, edge := range edges {
		graph[edge[0]][work[edge[0]]] = edge[1]
		work[edge[0]]++
		graph[edge[1]][work[edge[1]]] = edge[0]
		work[edge[1]]++
	}
	fmt.Println("work ", work)
	fmt.Println("graph ", graph)

	leafCnt, nextLeafCnt, numVertices := 0, 0, n
	for i := 0; i < n; i++ {
		// 度数为1的结点为叶子结点
		if work[i] == 1 {
			// 存储叶子结点下标
			work[leafCnt], leafCnt = i, leafCnt+1
		}
	}
	fmt.Println("work leafCnt ", work)
	for numVertices > 2 {
		nextLeafCnt = 0
		for i := 0; i < leafCnt; i++ {
			leaf := work[i]
			for _, index := range graph[leaf] {
				numEdges[index]--
				if numEdges[index] == 1 {
					level[nextLeafCnt], nextLeafCnt = index, nextLeafCnt+1
					fmt.Println("level ", level)
				}
			}
		}

		numVertices -= leafCnt
		work, level = level, work
		leafCnt = nextLeafCnt
	}
	if leafCnt == 2 {
		return []int{work[0], work[1]}
	}
	return []int{work[0]}
}

// func main()  {
//	edges := [][]int{{1,0},{1,2},{1,3}}
//	//numEdges := make([]int, 4)
//	//for _, edge := range edges {
//	//	numEdges[edge[0]]++
//	//	numEdges[edge[1]]++
//	//}
//	//fmt.Println(goodfindMinHeightTrees(4, edges))
//	fmt.Println(bfsfindMinHeightTrees(4, edges))
// }
