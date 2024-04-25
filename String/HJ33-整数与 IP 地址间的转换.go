package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hj33() {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		s := reader.Text()
		fmt.Println(convert(s))
	}
}

func convert(s string) string {
	// 截取字符串
	strs := strings.Split(s, ".")
	// 判断是整数还是IP地址
	if len(strs) == 4 {
		return ipToInt(strs)
	} else {
		return intToIp(s)
	}
}

func ipToInt(strs []string) string {
	// 拼接二进制数
	var ipInt uint32 = 0
	for _, v := range strs {
		num, _ := strconv.Atoi(v)
		ipInt = ipInt<<8 | uint32(num)
	}
	// 转换为十进制数
	return strconv.FormatUint(uint64(ipInt), 10)
}

func intToIp(s string) string {
	// 转换为ip地址
	ipInt, _ := strconv.ParseUint(s, 10, 32)
	return fmt.Sprintf("%d.%d.%d.%d",
		// ip的最高位能保证是8位，不用&0xff
		byte(ipInt>>24),
		byte(ipInt>>16&0xff),
		byte(ipInt>>8&0xff),
		byte(ipInt&0xff))
}

// func main() {
// 	hj33()
// }
