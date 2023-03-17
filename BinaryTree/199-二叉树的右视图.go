package main

/*199-二叉树的右视图*/
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil { // 防止为空
		return res
	}

	currentLevel := []*TreeNode{root}

	for len(currentLevel) > 0 {
		nextLevel := make([]*TreeNode, 0)
		var val []int
		for _, node := range currentLevel {
			// append node value
			val = append(val, node.Val)
			// append next level
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
		}
		currentLevel = nextLevel
		res = append(res, val[0])
	}
	return res
}
