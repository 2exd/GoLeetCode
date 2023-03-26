package main

import "math"

/*530.二叉搜索树的最小绝对差 easy*/
func getMinimumDifferenceMy(root *TreeNode) int {
	var res []int
	findMIn(root, &res)
	min := 1000000 // 一个比较大的值
	for i := 1; i < len(res); i++ {
		tempValue := res[i] - res[i-1]
		if tempValue < min {
			min = tempValue
		}
	}
	return min
}

// 中序遍历
func findMIn(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	findMIn(root.Left, res)
	*res = append(*res, root.Val)
	findMIn(root.Right, res)
}

// 中序遍历的同时计算最小值
func getMinimumDifference(root *TreeNode) int {
	// 保留前一个节点的指针
	var prev *TreeNode
	// 定义一个比较大的值
	min := math.MaxInt64
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev != nil && node.Val-prev.Val < min {
			min = node.Val - prev.Val
		}
		prev = node
		travel(node.Right)
	}
	travel(root)
	return min
}
