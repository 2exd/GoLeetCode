package main

/*236. 二叉树的最近公共祖先 medium*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// check if is leaf node
	if root == nil {
		return nil
	}

	// find target return
	if p == root || q == root {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// find target node q,p
	// all find
	if left != nil && right != nil {
		return root
	}
	// find on the left subtree
	if left != nil {
		return left
	}

	// find on the right subtree
	if right != nil {
		return right
	}

	// didn't find
	return nil
}
