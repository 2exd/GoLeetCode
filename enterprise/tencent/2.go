package main

var (
	path2 []bool
	res2  []string
	n     int
	strs  []string
)

// func main() {
// 	fmt.Scan(&n)
// 	strs = make([]string, n)
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&strs[i])
// 	}
// 	path2 = make([]bool, 128)
// 	// fmt.Print("2")
// 	var tmp string
// 	dfs2(tmp, 0)
// 	fmt.Print(len(res2))
// }

func dfs2(strNow string, start int) {
	if start == n {
		res2 = append(res2, strNow)
		return
	}
	str := strs[start]
	for i := 0; i < len(str); i++ {
		charId := str[i]
		if !path2[charId] {
			path2[charId] = true
			tmp := strNow + string(str[i])
			dfs2(tmp, start+1)
			path2[charId] = false
		}
	}
}
