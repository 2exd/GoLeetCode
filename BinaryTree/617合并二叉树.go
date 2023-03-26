package main

/*617. 合并二叉树 easy*/

// 递归
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	if root1 == nil && root2 == nil {
		return nil
	}

	// 构造root
	root := &TreeNode{Val: root1.Val + root2.Val,
		Left:  mergeTrees(root1.Left, root2.Left),
		Right: mergeTrees(root1.Right, root2.Right)}
	return root
}

// 迭代版本
func mergeTreesIteration(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	queue := make([]*TreeNode, 0)
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	queue = append(queue, root1)
	queue = append(queue, root2)

	for size := len(queue); size > 0; size = len(queue) {
		node1 := queue[0]
		queue = queue[1:]
		node2 := queue[0]
		queue = queue[1:]
		node1.Val += node2.Val
		// 左子树都不为空
		if node1.Left != nil && node2.Left != nil {
			queue = append(queue, node1.Left)
			queue = append(queue, node2.Left)
		}
		// 右子树都不为空
		if node1.Right != nil && node2.Right != nil {
			queue = append(queue, node1.Right)
			queue = append(queue, node2.Right)
		}
		// 树 1 的左子树为 nil，直接接上树 2 的左子树
		if node1.Left == nil {
			node1.Left = node2.Left
		}
		// 树 1 的右子树为 nil，直接接上树 2 的右子树
		if node1.Right == nil {
			node1.Right = node2.Right
		}
	}
	return root1
}
