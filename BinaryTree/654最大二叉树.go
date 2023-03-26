package main

/*654. 最大二叉树 medium*/
func findMaxIndex(order []int) (index int) {
	maxNum := -1
	index = 0
	for i := 0; i < len(order); i++ {
		if maxNum < order[i] {
			maxNum = order[i]
			index = i
		}
	}
	return index
}

// 递归
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}

	left := findMaxIndex(nums)
	// 构造root
	root := &TreeNode{Val: nums[left],
		Left:  constructMaximumBinaryTree(nums[:left]), // 将后续遍历一分为二，左边为左子树，右边为右子树
		Right: constructMaximumBinaryTree(nums[left+1:])}
	return root
}
