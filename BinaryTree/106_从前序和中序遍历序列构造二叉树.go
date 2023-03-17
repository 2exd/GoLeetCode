package main

/*106. 从中序与后序遍历序列构造二叉树 medium*/

// 递归
/**
来看一下一共分几步：
第一步：如果数组大小为零的话，说明是空节点了。
第二步：如果不为空，那么取后序数组最后一个元素作为节点元素。
第三步：找到后序数组最后一个元素在中序数组的位置，作为切割点
第四步：切割中序数组，切成中序左数组和中序右数组 （顺序别搞反了，一定是先切中序数组）
第五步：切割后序数组，切成后序左数组和后序右数组
第六步：递归处理左区间和右区间
*/
func buildTree106(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 || len(postorder) < 1 {
		return nil
	}
	// 先找到根节点（后续遍历的最后一个就是根节点）
	nodeValue := postorder[len(postorder)-1]
	// 从中序遍历中找到一分为二的点，左边为左子树，右边为右子树
	left := findRootIndex106(inorder, nodeValue)
	// 构造root
	root := &TreeNode{Val: nodeValue,
		Left:  buildTree106(inorder[:left], postorder[:left]), // 将后续遍历一分为二，左边为左子树，右边为右子树
		Right: buildTree106(inorder[left+1:], postorder[left:len(postorder)-1])}
	return root
}
func findRootIndex106(inorder []int, target int) (index int) {
	for i := 0; i < len(inorder); i++ {
		if target == inorder[i] {
			return i
		}
	}
	return -1
}
