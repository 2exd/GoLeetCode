package main

/*100. 相同的树 easy*/

// 递归

// 递归
func defs100(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return defs100(left.Left, right.Left) && defs100(left.Right, right.Right)
}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return defs100(p, q)
}

// 迭代
func isSameTreeStack(p *TreeNode, q *TreeNode) bool {
	var queue []*TreeNode
	queue = append(queue, p, q)
	for len(queue) > 0 {
		// 成对拿出来比较
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Left, left.Right, right.Right)
	}
	return true
}
