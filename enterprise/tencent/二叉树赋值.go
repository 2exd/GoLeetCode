package main

/*
二叉树共有 n 个节点，将所有节点赋值为 1~n 的正整数，且没有两个节点的值相等。
需要满足：奇数层的权值和偶数层的权值和之差的绝对值不超过1
*/

/*
思路是先统计偶数层一共要多少个数
然后将 1 - n 分成两个部分
两个部分的和差 1 就行了
*/

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func assignment(root *TreeNode) *TreeNode {
	evenCnt, nodeCnt, sum := 0, 0, 0

	que := make([]*TreeNode, 0)
	if root != nil {
		que = append(que, root)
	}
	layer := 0

	for len(que) > 0 {
		length := len(que)
		for i := 0; i < length; i++ {
			cur := que[0]
			que = que[1:]
			nodeCnt++
			if layer%2 == 0 {
				evenCnt++
			}
			if cur.left != nil {
				que = append(que, cur.left)
			}
			if cur.right != nil {
				que = append(que, cur.right)
			}
		}
		layer++
	}
	var target int
	sumTemp := calSum(nodeCnt)
	if sumTemp%2 == 0 {
		target = sumTemp / 2
	} else {
		target = sumTemp/2 + 1
	}

	eoHash := make([]int, nodeCnt+1)
	for i := 1; i <= evenCnt-1; i++ {
		eoHash[i] = 1
		sum += i
	}
	if eoHash[target-sum] == 1 {
		return nil
	}
	eoHash[target-sum] = 1

	even, odd := make([]int, 0), make([]int, 0)
	for i := 1; i < nodeCnt; i++ {
		if eoHash[i] == 1 {
			even = append(even, i)
		} else {
			odd = append(odd, i)
		}
	}

	que = append(que, root)
	e, o := 0, 0
	for len(que) > 0 {
		size := len(que)
		for i := 0; i < size; i++ {
			cur := que[0]
			que = que[1:]
			cur.val = even[e]
			e++
			if cur.left != nil {
				que = append(que, cur.left)
			}
			if cur.right != nil {
				que = append(que, cur.right)
			}
		}
		size = len(que)
		for i := 0; i < size; i++ {
			cur := que[0]
			que = que[1:]
			cur.val = odd[o]
			o++
			if cur.left != nil {
				que = append(que, cur.left)
			}
			if cur.right != nil {
				que = append(que, cur.right)
			}
		}
	}
	return root
}

func calSum(cnt int) int {
	sum := 0
	for i := 1; i <= cnt; i++ {
		sum += i
	}
	return sum
}

/*func main() {

}
*/
