package main

import "container/list"

/*572. 另一棵树的子树 easy*/

// 递归
func defs572(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return defs572(left.Left, right.Left) && defs572(left.Right, right.Right)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return defs572(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSubtreeStack(root *TreeNode, subRoot *TreeNode) bool {

	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)

		if defs572(node, subRoot) {
			return true
		}
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return false
}
