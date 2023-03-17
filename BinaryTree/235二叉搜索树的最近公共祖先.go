package main

/*235. 二叉搜索树的最近公共祖先 easy*/

/*
其实只要从上到下遍历的时候，cur节点是数值在[p, q]区间中则说明该节点cur就是最近公共祖先了。
*/
//利用BSL的性质（前序遍历有序）
func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val { //当前节点的值大于给定的值，则说明满足条件的在左边
		return lowestCommonAncestorBST(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val { //当前节点的值小于各点的值，则说明满足条件的在右边
		return lowestCommonAncestorBST(root.Right, p, q)
	} else {
		return root
	} //当前节点的值在给定值的中间（或者等于），即为最深的祖先
}
