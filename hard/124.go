package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*方法一：递归
首先计算每个结点的最大贡献值(具体而言，就是在以该节点为根节点的子树中寻找以该节点为起点的一条路径，使得该路径上的节点值之和最大)
该函数的计算如下。
空节点的最大贡献值等于 0。
非空节点的最大贡献值等于节点值与其子节点中的最大贡献值之和（对于叶节点而言，最大贡献值等于节点值）

然后得到二叉树的最大路径和
如果子节点的最大贡献值为正，则计入该节点的最大路径和，否则不计入该节点的最大路径和
维护一个全局变量 maxSum 存储最大路径和，在递归过程中更新 maxSum 的值，最后得到的 maxSum 的值即为二叉树中的最大路径和
*/
func maxPathSum(root *TreeNode) int {
	// -2147483648
	maxSum := math.MinInt32

	// 计算node最大贡献值
	var maxNodeSum func(*TreeNode) int
	maxNodeSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算子节点的最大贡献值
		leftNodeSum := max(maxNodeSum(node.Left), 0)
		rightNodeSum := max(maxNodeSum(node.Right), 0)

		// 计算当前结点总值
		currentSum := node.Val + leftNodeSum + rightNodeSum

		// 更新当前贡献总值
		maxSum = max(maxSum, currentSum)

		// 返回最大结点贡献值
		return node.Val + max(leftNodeSum, rightNodeSum)
	}
	maxNodeSum(root)

	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*复杂度分析
时间复杂度：O(N)，其中 N 是二叉树中的节点个数。对每个节点访问不超过 2 次。
空间复杂度：O(N)，其中 N 是二叉树中的节点个数。空间复杂度主要取决于递归调用层数，最大层数等于二叉树的高度，最坏情况下，二叉树的高度等于二叉树中的节点个数。*/
