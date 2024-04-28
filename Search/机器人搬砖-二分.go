package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func main() {
//
// }

func findMinPower() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// 将输入的字符串分割成整数数组
	bricks := parseInput(input)

	// 如果仓库数量超过8个，则无法在8小时内完成，直接返回-1
	if len(bricks) > 8 {
		fmt.Println(-1)
		return
	}

	// 使用二分查找算法找到每小时最少需要充多少个能量格
	result := minEnergyPerHour(bricks)
	fmt.Println(result)
}

// parseInput 将输入字符串转换为整数切片
func parseInput(input string) []int {
	parts := strings.Fields(input)
	bricks := make([]int, len(parts))
	for i, p := range parts {
		bricks[i], _ = strconv.Atoi(p)
	}
	return bricks
}

// minEnergyPerHour 使用二分查找算法确定机器人每小时最少需要多少能量格
func minEnergyPerHour(bricks []int) int {
	l, r := 1, max1(bricks)
	for l < r {
		m := (l + r) / 2
		if canCompleteInEightHours(bricks, m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return r
}

// max 返回切片中的最大值
func max1(arr []int) int {
	maxVal := arr[0]
	for _, val := range arr[1:] {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

// canCompleteInEightHours 检查给定的能量格数是否足以在8小时内完成搬砖任务
func canCompleteInEightHours(bricks []int, power int) bool {
	totalHours := 0
	for _, brick := range bricks {
		totalHours += (brick + power - 1) / power // 向上取整计算所需小时数
	}
	return totalHours <= 8
}
