package main

import (
	"sort"
)

/*435. 无重叠区间*/
func eraseOverlapIntervals(intervals [][]int) int {
	// 	order by start station asc
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := 0
	for i := 1; i < len(intervals); i++ {
		// if x end position > x+1 start position, no need to remove
		if intervals[i-1][1] > intervals[i][0] {
			result++
			// update smallest end position
			intervals[i][1] = min435(intervals[i-1][1], intervals[i][1])
		}
	}
	return result
}

func min435(a, b int) int {
	if a > b {
		return b
	}
	return a
}
