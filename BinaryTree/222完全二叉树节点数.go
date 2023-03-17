package main

/*222. 完全二叉树的节点个数 medium*/
// 层序遍历
/*
时间复杂度：O(n)
空间复杂度：O(n)
*/
func countNodesErgodic(root *TreeNode) int {
	nums := 0
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
		nums++
	}
	for l := len(queue); l > 0; {
		for ; l > 0; l-- {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
				nums++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nums++
			}
			queue = queue[1:]
		}
		l = len(queue)
	}
	return nums
}

/*
时间复杂度：O(n)
空间复杂度：O(log n)，算上了递归系统栈占用的空间
*/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	if root.Right != nil {
		res += countNodes(root.Right)
	}
	if root.Left != nil {
		res += countNodes(root.Left)
	}
	return res
}
