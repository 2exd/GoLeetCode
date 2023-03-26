package main

/*321. 拼接最大数*/

/*单调栈
为了找到长度为 k 的最大数，需要从两个数组中分别选出最大的子序列，这两个子序列的长度之和为 k，然后将这两个子序列合并得到最大数。两个子序列的长度最小为 0，最大不能超过 k 且不能超过对应的数组长度
*/

/*假设我们从 nums1 中取了 k1 个，从 num2 中取了 k2 个，其中 k1 + k2 = k。而 k1 和 k2 这 两个子问题我们是会解决的。由于这两个子问题是相互独立的，因此我们只需要分别求解，然后将结果合并即可。
假如 k1 和 k2 个数字，已经取出来了。那么剩下要做的就是将这个长度分别为 k1 和 k2 的数字，合并成一个长度为 k 的数组合并成一个最大的数组。
那剩余的问题就是怎么在num中取出k个最大的数字并保持相对顺序不变呢？其实跟题目4：402. 移掉K位数字 是一样的操作，使用单调栈！
实际上这个过程有点类似归并排序中的治，而上面我们分别计算 num1 和 num2 的最大数的过程类似归并排序中的分*/

func maxSubsequence(a []int, k int) (s []int) {
	// 遍历 a
	for i, v := range a {
		// s长度大于0 且 当前子序列长度+a剩余长度>=目标子序列k长
		for len(s) > 0 && len(s)+len(a)-1-i >= k && v > s[len(s)-1] {
			// 当前元素比栈尾大就弹栈
			s = s[:len(s)-1]
		}
		if len(s) < k {
			s = append(s, v)
		}
	}
	return
}

func lexicographicalLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			// 比位的大小
			return a[i] < b[i]
		}
	}
	// 比总长度
	return len(a) < len(b)
}

func merge(a, b []int) []int {
	merged := make([]int, len(a)+len(b))
	for i := range merged {
		// a<b
		if lexicographicalLess(a, b) {
			// 弹栈首，进入merged
			merged[i], b = b[0], b[1:]
		} else {
			merged[i], a = a[0], a[1:]
		}
	}
	return merged
}

func maxNumber(nums1, nums2 []int, k int) (res []int) {
	start := 0
	if k > len(nums2) {
		start = k - len(nums2)
	}
	for i := start; i <= k && i <= len(nums1); i++ {
		s1 := maxSubsequence(nums1, i)
		s2 := maxSubsequence(nums2, k-i)
		merged := merge(s1, s2)
		// res<merged
		if lexicographicalLess(res, merged) {
			res = merged
		}
	}
	return
}

/*复杂度分析
时间复杂度：O(k(m+n+k^2))，其中 m 和 n 分别是数组 nums1 和 nums2的长度，k 是拼接最大数的长度。
两个子序列的长度之和为 k，最多有 k 种不同的长度组合。对于每一种长度组合，需要首先得到两个最大子序列，然后进行合并。
得到两个最大子序列的时间复杂度为线性，即 O(m+n)。
合并两个最大子序列，需要进行 k 次合并，每次合并需要进行比较，最坏情况下，比较的时间复杂度为 O(k)，因此合并操作的时间复杂度为 O(k^2)。
因此对于每一种长度组合，时间复杂度为 O(m+n+k^2)，总时间复杂度为 O(k(m+n+k^2))

空间复杂度：O(k)，其中 k 是拼接最大数的长度。每次从两个数组得到两个子序列，两个子序列的长度之和为 k。*/
