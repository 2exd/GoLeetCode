package main

/*226. 翻转二叉树 easy*/
// 只需要交换左右子结点
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left // 交换

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

func invertTreeStack(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			node.Left, node.Right = node.Right, node.Left // 交换
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	return root
}
