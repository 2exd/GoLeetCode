package main

func trapDoublePointer(height []int) (ans int) {
	for i := 0; i < len(height); i++ {
		// 首尾不接雨水
		if i == 0 || i == len(height)-1 {
			continue
		}
		// 记录左右最高柱子高度
		rHeight := height[i]
		lHeight := height[i]
		for r := i + 1; r < len(height); r++ {
			if height[r] > rHeight {
				rHeight = height[r]
			}
		}
		for l := i - 1; l >= 0; l-- {
			if height[l] > lHeight {
				lHeight = height[l]
			}
		}

		// 取最小柱子高度
		h := findMin42(lHeight, rHeight) - height[i]
		if h > 0 {
			ans += h
		}
	}
	return ans
}

func findMin42(x, y int) int {
	if x < y {
		return x
	}
	return y
}
