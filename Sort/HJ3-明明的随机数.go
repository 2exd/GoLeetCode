package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func hj3() {
	// 将随机数存储在数组中, 索引当作随机数值, 数组值当作是否存在
	data := make([]bool, 501)
	scanner := bufio.NewScanner(os.Stdin)
	// 读取数据范围
	scanner.Scan()
	cntText := scanner.Text()
	cnt, _ := strconv.Atoi(cntText)
	for cnt > 0 && scanner.Scan() {
		cnt--
		text := scanner.Text()
		num, err := strconv.Atoi(text)
		if err != nil {
			continue
		}
		// 标记随机数位置
		data[num] = true
	}
	// 有序数组
	for i := 1; i <= 500; i++ {
		if data[i] {
			fmt.Printf("%d\n", i)
		}
	}
}
