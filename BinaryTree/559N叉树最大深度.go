package main

/*559. N 叉树的最大深度 easy*/

func max559(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 递归
func maxDepth559Recursion(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	for i := 0; i < len(root.Children); i++ {
		depth = max(depth, maxDepth559Recursion(root.Children[i]))
	}
	return depth + 1
}

// 遍历
func maxDepth559(root *Node) int {
	levl := 0
	queue := make([]*Node, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		for ; l > 0; l-- {
			node := queue[0]
			for i := 0; i < len(node.Children); i++ {
				if node.Children[i] != nil {
					queue = append(queue, node.Children[i])
				}
			}

			queue = queue[1:]
		}
		levl++
		l = len(queue)
	}
	return levl
}

type Node struct {
	Val      int
	Children []*Node
}
