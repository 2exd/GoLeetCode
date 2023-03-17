package main

import "container/list"

/*513. 找树左下角的值 medium*/
func findBottomLeftValueIteraion(root *TreeNode) int {
	var res int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if i == 0 {
				res = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return res
}

var maxDeep int // 全局变量 深度
var value int   //全局变量 最终值
func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil { //需要提前判断一下（不要这个if的话提交结果会出错，但执行代码不会。防止这种情况出现，故先判断是否只有一个节点）
		return root.Val
	}
	findLeftValue(root, maxDeep)
	return value
}
func findLeftValue(root *TreeNode, deep int) {
	//最左边的值在左边
	if root.Left == nil && root.Right == nil {
		if deep > maxDeep {
			value = root.Val
			maxDeep = deep
		}
	}
	//递归
	if root.Left != nil {
		deep++
		findLeftValue(root.Left, deep)
		deep-- //回溯
	}
	if root.Right != nil {
		deep++
		findLeftValue(root.Right, deep)
		deep-- //回溯
	}
}
