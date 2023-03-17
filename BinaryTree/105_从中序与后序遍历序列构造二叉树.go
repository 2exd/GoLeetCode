package main

/*105. 从前序与中序遍历序列构造二叉树 medium*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}

	left := findRootIndex(preorder[0], inorder)

	root := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:left+1], inorder[:left]),
		Right: buildTree(preorder[left+1:], inorder[left+1:]),
	}

	return root
}

func findRootIndex(target int, inorder []int) int {
	for i, v := range inorder {
		if v == target {
			return i
		}
	}
	return -1
}
