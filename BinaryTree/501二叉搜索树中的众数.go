package main

/*501. 二叉搜索树中的众数 easy*/
func findModeMy(root *TreeNode) []int {
	mapNum := map[int]int{}
	var res []int
	findMost(root, &mapNum)
	max := 0 //一个比较大的值
	for _, value := range mapNum {
		if value > max {
			max = value
		}
	}
	for key, value := range mapNum {
		if value == max {
			res = append(res, key)
		}
	}
	return res
}

//中序遍历
func findMost(root *TreeNode, mapNum *map[int]int) {
	if root == nil {
		return
	}
	findMost(root.Left, mapNum)
	(*mapNum)[root.Val]++
	findMost(root.Right, mapNum)
}

/*计数法，不使用额外空间，利用二叉树性质，中序遍历*/
func findMode(root *TreeNode) []int {
	res := make([]int, 0)
	count := 1
	max := 1
	var prev *TreeNode
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev != nil && prev.Val == node.Val {
			count++
		} else {
			count = 1
		}
		if count >= max {
			if count > max && len(res) > 0 {
				res = []int{node.Val}
			} else {
				res = append(res, node.Val)
			}
			max = count
		}
		prev = node
		travel(node.Right)
	}
	travel(root)
	return res
}
