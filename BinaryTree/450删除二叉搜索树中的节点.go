package main

/*450.删除二叉搜索树中的节点 medium*/
/*
删除节点有以下五种情况：
第一种情况：没找到删除的节点，遍历到空节点直接返回了
第二种情况：左右孩子都为空（叶子节点），直接删除节点， 返回NULL为根节点
第三种情况：删除节点的左孩子为空，右孩子不为空，删除节点，右孩子补位，返回右孩子为根节点
第四种情况：删除节点的右孩子为空，左孩子不为空，删除节点，左孩子补位，返回左孩子为根节点
*****第五种情况：左右孩子节点都不为空，则将删除节点的左子树头结点（左孩子）放到删除节点的右子树的最左面节点的左孩子上，返回删除节点右孩子为新的根节点。
*/

// 递归版本
func deleteNode(root *TreeNode, key int) *TreeNode {
	// 1
	if root == nil {
		return nil
	}
	// 遍历
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	// 3
	if root.Right == nil {
		return root.Left
	}
	// 4
	if root.Left == nil {
		return root.Right
	}
	// minnode := root.Right
	// for minnode.Left != nil {
	//	minnode = minnode.Left
	// }
	// root.Val = minnode.Val
	root.Val = root.Right.Val
	root.Right = deleteNode1(root.Right)
	return root
}

func deleteNode1(root *TreeNode) *TreeNode {
	if root.Left == nil {
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	root.Left = deleteNode1(root.Left)
	return root
}

// 迭代版本
func deleteOneNodeIteration(target *TreeNode) *TreeNode {
	if target == nil {
		return target
	}
	if target.Right == nil {
		return target.Left
	}
	cur := target.Right
	for cur.Left != nil {
		cur = cur.Left
	}
	cur.Left = target.Left
	return target.Right
}
func deleteNodeIteration(root *TreeNode, key int) *TreeNode {
	// 特殊情况处理
	if root == nil {
		return root
	}
	cur := root
	var pre *TreeNode
	for cur != nil {
		if cur.Val == key {
			break
		}
		pre = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if pre == nil {
		return deleteOneNodeIteration(cur)
	}
	// pre 要知道是删除左孩子还有右孩子
	if pre.Left != nil && pre.Left.Val == key {
		pre.Left = deleteOneNodeIteration(cur)
	}
	if pre.Right != nil && pre.Right.Val == key {
		pre.Right = deleteOneNodeIteration(cur)
	}
	return root
}
