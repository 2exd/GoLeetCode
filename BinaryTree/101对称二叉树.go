package main

/*101. 对称二叉树 easy*/

// 递归
func isSymmetric(root *TreeNode) bool {
	return dfs101(root.Left, root.Right)
}

func dfs101(left *TreeNode, right *TreeNode) bool {
	// if both are nil
	if left == nil && right == nil {
		return true
	}
	// compare the pointers
	if left == nil || right == nil {
		return false
	}
	// don't forget to compare values'
	if left.Val != right.Val {
		return false
	}
	// notice the sequence!
	return dfs101(left.Left, right.Right) && dfs101(right.Left, left.Right)
}

func isSymmetricQueue(root *TreeNode) bool {
	currentLevel := make([]*TreeNode, 0)
	if root != nil {
		currentLevel = append(currentLevel, root.Left, root.Right)
	}
	for len(currentLevel) > 0 {
		left := currentLevel[0]
		right := currentLevel[1]
		currentLevel = currentLevel[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		// notice the sequence!
		currentLevel = append(currentLevel, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}
