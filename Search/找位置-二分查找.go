package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findSeat() {
	scanner := bufio.NewScanner(os.Stdin)

	// 读取第一行输入，已经排好的学生学号
	scanner.Scan()
	studentNumbers := strings.Fields(scanner.Text())

	// 将字符串数组转换为整数数组
	students := make([]int, len(studentNumbers))
	for i, v := range studentNumbers {
		students[i], _ = strconv.Atoi(v)
	}

	// 读取第二行，小明的学号
	scanner.Scan()
	xiaomingNumber, _ := strconv.Atoi(scanner.Text())

	// 使用二分搜索找到小明应该插入的位置
	position := findInsertPosition(students, xiaomingNumber)

	// 输出小明应该站的位置，位置是从1开始的
	fmt.Println(position + 1)
}

// findInsertPosition 使用二分搜索算法确定插入位置
func findInsertPosition(arr []int, target int) int {
	low, high := 0, len(arr)
	for low < high {
		mid := low + (high-low)/2
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
