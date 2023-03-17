package main

/*103-二叉树的锯齿形层序遍历*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil { // 防止为空
		return res
	}
	forward := false
	currentLevel := []*TreeNode{root}

	for len(currentLevel) > 0 {
		nextLevel := make([]*TreeNode, 0)
		var val []int
		for _, node := range currentLevel {
			// append node value
			val = append(val, node.Val)
			// append next level
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		if !forward {
			for i, n := 0, len(val); i < n/2; i++ {
				val[i], val[n-1-i] = val[n-1-i], val[i]
			}
		}
		forward = !forward
		currentLevel = nextLevel
		res = append(res, val)
	}
	return res
}
