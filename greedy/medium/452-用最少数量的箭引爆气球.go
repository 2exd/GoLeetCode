package main

import "sort"

func findMinArrowShots(points [][]int) int {
	// number of arrows
	var res = 1
	// order by X start position
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	for i := 1; i < len(points); i++ {
		// if x end position < x+1 start position, must not coincide, arrows++;
		if points[i-1][1] < points[i][0] {
			res++
		} else {
			// update balloons coincident smallest end position, reserved for next step
			points[i][1] = min452(points[i-1][1], points[i][1])
		}
	}
	return res
}
func min452(a, b int) int {
	if a > b {
		return b
	}
	return a
}
