package main

/*
88-合并有序数组
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	var temp []int
	for i < m && j < n {
		if nums1[i] >= nums2[j] {
			temp = append(temp, nums2[j])
			j++
		} else {
			temp = append(temp, nums1[i])
			i++
		}
	}

	if i < m {
		temp = append(temp, nums1[i:m]...)
	}

	if j < n {
		temp = append(temp, nums2[j:n]...)
	}

	copy(nums1, temp)
}

//
// func main() {
// 	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
// }
