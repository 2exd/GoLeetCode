package main

import (
	"math/rand"
	"time"
)

func sortArray(nums []int) []int {
	rand.Seed(time.Now().UnixNano())
	quickSort912(nums)
	return nums
}

func quickSort912(nums []int) {
	length := len(nums)
	rIndex := rand.Int()%length - 1
	nums[0], nums[rIndex] = nums[rIndex], nums[0]
	mid := nums[0]
	l, r := 0, length-1
	allsame := true
	for l < r {
		// find elements smaller than mid, put them on the left side
		for l < r {
			if nums[r] > mid {
				allsame = false
				r--
			} else if nums[r] == mid {
				r--
			} else {
				nums[l] = nums[r]
				break
			}
		}
		// find elements bigger than mid, put them on the right side
		for l < r {
			if nums[l] < mid {
				allsame = false
				l++
			} else if nums[l] == mid {
				l++
			} else {
				nums[r] = nums[l]
				break
			}
		}
	}

	nums[l] = mid

	// 保证:
	// - nums[x] < nums[i], x < i   // - nums[x] > nums[i], x > i
	if allsame {
		return
	}

	quickSort912(nums[:l])
	quickSort912(nums[l+1:])
}

// func main() {
// 	quickSort912([]int{5, 1, 1, 2, 0, 0})
// }
