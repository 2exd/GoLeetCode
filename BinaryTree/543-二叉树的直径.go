package main

/*543-二叉树的直径*/
func diameterOfBinaryTree(root *TreeNode) int {
	return dfs(root)[0]
}

func dfs(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{-1, -1}
	}
	l := dfs(root.Left)
	r := dfs(root.Right)
	a := max(l[1]+r[1]+2, max(l[0], r[0]))
	b := max(l[1], r[1]) + 1
	return [2]int{a, b}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
