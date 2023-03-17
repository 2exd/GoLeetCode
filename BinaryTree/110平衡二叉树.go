package main

/*110. 平衡二叉树 easy*/
// 递归
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	lDepth := findMaxDepth(root.Left)
	rDepth := findMaxDepth(root.Right)
	return abs110(lDepth-rDepth) <= 1
}

func findMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return findMax110(findMaxDepth(root.Left), findMaxDepth(root.Right)) + 1
}

func findMax110(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func abs110(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
