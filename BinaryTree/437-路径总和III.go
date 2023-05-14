package main

/*
437-路径总和III
*/

// dfs
func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	val := root.Val
	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return
}

func pathSum437(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum437(root.Left, targetSum)
	res += pathSum437(root.Right, targetSum)
	return res
}

// 前缀和
func pathSumPrefixSum(root *TreeNode, targetSum int) (ans int) {
	preSum := map[int64]int{0: 1}
	var dfs func(*TreeNode, int64)
	dfs = func(node *TreeNode, curr int64) {
		if node == nil {
			return
		}
		curr += int64(node.Val)
		ans += preSum[curr-int64(targetSum)]
		preSum[curr]++
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		preSum[curr]--
		return
	}
	dfs(root, 0)
	return
}
