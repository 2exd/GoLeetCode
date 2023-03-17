package main

/*111. 二叉树的最小深度 easy*/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 递归
func minDepthRecursion(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right != nil {
		return min(minDepthRecursion(root.Left), minDepthRecursion(root.Right)) + 1
	} else if root.Left == nil && root.Right != nil {
		return minDepthRecursion(root.Right) + 1
	} else if root.Left != nil && root.Right == nil {
		return minDepthRecursion(root.Left) + 1
	} else {
		return 0
	}
}

// 层序遍历
func minDepth(root *TreeNode) int {
	dep := 0
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		dep++
		for ; l > 0; l-- {
			node := queue[0]
			// 如果遇到叶子结点，就直接返回
			if node.Left == nil && node.Right == nil {
				return dep
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		l = len(queue)
	}
	return dep
}
