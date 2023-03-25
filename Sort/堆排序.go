package main

func heapSortArray(nums []int) []int {
	// 堆排序-大根堆，升序排序，基于比较交换的不稳定算法，时间O(nlogn)，空间O(1)-迭代建堆
	// 遍历元素时间O(n)，堆化时间O(logn)，开始建堆次数多些，后面次数少
	// 主要思路：
	// 1.建堆，从非叶子节点开始依次堆化，注意逆序，从下往上堆化
	// 建堆流程：父节点与子节点比较，子节点大则交换父子节点，父节点索引更新为子节点，循环操作
	// 2.尾部遍历操作，弹出元素，再次堆化
	// 弹出元素排序流程：从最后节点开始，交换头尾元素，由于弹出，end--，再次对剩余数组元素建堆，循环操作
	// 建堆函数，堆化
	var heapify func(nums []int, root, end int)
	heapify = func(nums []int, root, end int) {
		// 大顶堆堆化，堆顶值小一直下沉
		for {
			// 左孩子节点索引
			child := root*2 + 1
			// 越界跳出
			if child > end {
				return
			}
			// 比较左右孩子，取大值，否则child不用++
			if child < end && nums[child] <= nums[child+1] {
				child++
			}
			// 如果父节点已经大于左右孩子大值，已堆化.
			if nums[root] > nums[child] {
				return
			}
			// 孩子节点大值上冒
			nums[root], nums[child] = nums[child], nums[root]
			// 更新父节点到子节点，继续往下比较，不断下沉
			root = child
		}
	}
	end := len(nums) - 1
	// 从最后一个非叶子节点开始堆化
	for i := end / 2; i >= 0; i-- {
		heapify(nums, i, end)
	}
	// 依次弹出元素，然后再堆化，相当于依次把最大值放入尾部
	for i := end; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		end--
		heapify(nums, 0, end)
	}
	return nums
}
