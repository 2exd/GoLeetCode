package main

/*104. 二叉树的最大深度 easy*/
func maxDepth(root *TreeNode) int {
	level := 0
	var curLevel []*TreeNode
	if root != nil {
		curLevel = append(curLevel, root)
	}
	for len(curLevel) > 0 {
		nextLevel := make([]*TreeNode, 0)
		for _, node := range curLevel {
			if node == nil {
				continue
			}
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		level++
		curLevel = nextLevel
	}
	return level
}
