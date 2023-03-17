package main

/*112. 路径总和 easy*/

// 递归
func hasPathSumMy(root *TreeNode, targetSum int) bool {
	res := make([]int, 0)
	var travel func(node *TreeNode, s int)
	travel = func(node *TreeNode, s int) {
		if node == nil {
			return
		}
		// 遇到叶子结点就加入
		if node.Left == nil && node.Right == nil {
			v := s + node.Val
			res = append(res, v)
			return
		}
		// 前序遍历
		s = s + node.Val
		if node.Left != nil {
			travel(node.Left, s)
		}
		if node.Right != nil {
			travel(node.Right, s)
		}
	}
	travel(root, 0)
	for i := 0; i < len(res); i++ {
		if targetSum == res[i] {
			return true
		}
	}
	return false
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val                                        // 将targetSum在遍历每层的时候都减去本层节点的值
	if root.Left == nil && root.Right == nil && targetSum == 0 { // 如果在叶子结点剩余的targetSum为0, 则正好就是符合的结果
		return true
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum) // 否则递归找
}
