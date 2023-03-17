package main

/*102. 二叉树的层序遍历 medium*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
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
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}

		}
		currentLevel = nextLevel
		res = append(res, val)
	}

	return res
}
