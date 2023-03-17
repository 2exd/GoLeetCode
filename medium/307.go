package main

/*307. 区域和检索 - 数组可修改*/

/*
分块处理
*/

//type NumArray struct {
//	nums, sums []int // sums[i] 表示第 i 个块的元素和
//	size       int   // 块的大小
//}
//
//func Constructor(nums []int) NumArray {
//	n := len(nums)
//	size := int(math.Sqrt(float64(n)))
//	sums := make([]int, (n+size-1)/size) // n/size 向上取整
//	for i, num := range nums {
//		sums[i/size] += num
//	}
//	return NumArray{nums, sums, size}
//}
//
//func (na *NumArray) Update(index, val int) {
//	// 更新块的sum
//	na.sums[index/na.size] += val - na.nums[index]
//	na.nums[index] = val
//}
//
//func (na *NumArray) SumRange(left, right int) (ans int) {
//	size := na.size
//	b1, b2 := left/size, right/size
//	if b1 == b2 { // 区间 [left, right] 在同一块中
//		for i := left; i <= right; i++ {
//			ans += na.nums[i]
//		}
//		return
//	}
//	for i := left; i < (b1+1)*size; i++ {
//		ans += na.nums[i]
//	}
//	for i := b1 + 1; i < b2; i++ {
//		ans += na.sums[i]
//	}
//	for i := b2 * size; i <= right; i++ {
//		ans += na.nums[i]
//	}
//	return
//}
/*复杂度分析
时间复杂度：构造函数为 O(n)，update 函数为 O(1)，sumRange 函数为 O(\sqrt n)，
其中 n 为数组 nums 的大小。对于 sumRange 函数，我们最多遍历两个块以及 sum 数组，因此时间复杂度为 O(\sqrt n)。
空间复杂度：O(\sqrt n)。保存 sum 数组需要 O(\sqrt n)的空间。*/


/*线段树*/

//type NumArray []int
//
//func Constructor(nums []int) NumArray {
//	n := len(nums)
//	seg := make(NumArray, n*4)
//	seg.build(nums, 0, 0, n-1)
//	return seg
//}
//
//
//// 构建完全二叉树
//func (seg NumArray) build(nums []int, node, s, e int) {
//	// 叶子结点, 为单位区间
//	if s == e {
//		seg[node] = nums[s]
//		return
//	}
//	// 计算区间中点
//	m := s + (e-s)/2
//	seg.build(nums, node*2+1, s, m)// 左孩子区间为node*2+1
//	seg.build(nums, node*2+2, m+1, e)// 右孩子区间为node*2+2
//	seg[node] = seg[node*2+1] + seg[node*2+2]
//}
//
//func (seg NumArray) change(index, val, node, s, e int) {
//	if s == e {
//		seg[node] = val// 修改
//		return
//	}
//	m := s + (e-s)/2
//	if index <= m {
//		seg.change(index, val, node*2+1, s, m)
//	} else {
//		seg.change(index, val, node*2+2, m+1, e)
//	}
//	seg[node] = seg[node*2+1] + seg[node*2+2]// 更新
//}
//
//func (seg NumArray) range_(left, right, node, s, e int) int {
//	if left == s && right == e {
//		return seg[node]
//	}
//	m := s + (e-s)/2
//	if right <= m {
//		return seg.range_(left, right, node*2+1, s, m)
//	}
//	if left > m {
//		return seg.range_(left, right, node*2+2, m+1, e)
//	}
//	return seg.range_(left, m, node*2+1, s, m) + seg.range_(m+1, right, node*2+2, m+1, e)
//}
//
//func (seg NumArray) Update(index, val int) {
//	seg.change(index, val, 0, 0, len(seg)/4-1)
//}
//
//func (seg NumArray) SumRange(left, right int) int {
//	return seg.range_(left, right, 0, 0, len(seg)/4-1)
//}

/*复杂度分析
时间复杂度：
构造函数：O(n)，其中 n 是数组 nums 的大小。二叉树的高度不超过 ⌈logn⌉+1，那么segmentTree 的大小不超过 2 ^ (⌈logn⌉+1) −1≤4n，所以build 的时间复杂度为 O(n)。
update 函数：O(logn)。因为树的高度不超过 ⌈logn⌉+1，所以涉及更新的结点数不超过 ⌈logn⌉+1。
sumRange 函数：O(logn)。每层结点最多访问四个，总共访问的结点数不超过4×(⌈logn⌉+1)。

空间复杂度：O(n)。保存 \segmentTree 需要 O(n) 的空间。*/

/*树状数组*/

type NumArray struct {
	nums, tree []int
}

func Constructor(nums []int) NumArray {
	// bit tree第一个元素0不用
	tree := make([]int, len(nums)+1)
	na := NumArray{nums, tree}
	for i, num := range nums {
		na.add(i+1, num)
	}
	return na
}

func (na *NumArray) add(index, val int) {
	// lowbit操作 index & -index 求x最低位的1
	for ; index < len(na.tree); index += index & -index {
		// 存储sum值
		na.tree[index] += val
	}
}

// 计算某点sum值
func (na *NumArray) prefixSum(index int) (sum int) {
	for ; index > 0; index &= index - 1 {
		sum += na.tree[index]
	}
	return
}

func (na *NumArray) Update(index, val int) {
	na.add(index+1, val-na.nums[index])// 更新sum值
	na.nums[index] = val
}


func (na *NumArray) SumRange(left, right int) int {
	// 到right的sum - 到left-1的sum
	return na.prefixSum(right+1) - na.prefixSum(left)
}