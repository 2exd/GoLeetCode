package main

/*
题干
	小红拿到了一个字符串，她想知道这个字符串能否通过重新排列 组成"Baidu"字符串?
	注:必须大小写完全相同。
	共有t组询问。
输入描述
	第一行输入一个正整数t，代表询问次数。
	接下来的t行，每行输入一个仅包含英文字母的字符串。
	所有字符串的长度之和保证不超过200000。
输出描述
	对于每次询问，输出一行答案。如果可以通过重新排列组成"Baidu",则输出"Yes"， 否则输出"No"。
*/
/*
解题思路
	对于每个询问，读取一个字符串并调用 Form 函数来检查是否可以通过重新排列组成"Baidu"字符串。
	如果返回true，则输出"Yes"，否则输出"No"。
	在 Form 函数中，使用一个 Map 来统计输入字符串中每个字符出现的次数。
	然后检查是否有足够的字符可以组成"Baidu"。
*/

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	t, _ := strconv.Atoi(scanner.Text())
// 	for i := 0; i < t; i++ {
// 		scanner.Scan()
// 		s := scanner.Text()
// 		if Form(s) {
// 			fmt.Println("Yes")
// 		} else {
// 			fmt.Println("No")
// 		}
// 	}
// }

func Form(s string) bool {
	freq := make(map[rune]int)
	for _, c := range s {
		freq[c]++
	}
	if freq['B'] == 1 && freq['a'] == 1 && freq['i'] == 1 && freq['d'] == 1 && freq['u'] == 1 {
		return true
	}
	return false
}
