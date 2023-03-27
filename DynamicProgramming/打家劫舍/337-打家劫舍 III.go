package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

/*337-打家劫舍 III*/
func robIII(root *TreeNode) int {
	res := robTree(root)
	return max337(res[0], res[1])
}

func max337(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robTree(cur *TreeNode) []int {
	if cur == nil {
		return []int{0, 0}
	}
	// 后序遍历
	left := robTree(cur.Left)
	right := robTree(cur.Right)

	// 考虑去偷当前的屋子
	robCur := cur.Val + left[0] + right[0]
	// 考虑不去偷当前的屋子
	notRobCur := max337(left[0], left[1]) + max337(right[0], right[1])

	// 注意顺序：0:不偷，1:去偷
	return []int{notRobCur, robCur}
}
