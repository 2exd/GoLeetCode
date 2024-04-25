package main

import "fmt"

// 写出一个程序，接受一个十六进制的数，输出该数值的十进制表示。
// 31  −1

func change(x int) {
	var num int
	for {
		_, err := fmt.Scanf("0x%x", &num)
		if err != nil {
			return
		}
		fmt.Println(num)
	}
}
