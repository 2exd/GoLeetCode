package main

/*383. 赎金信 easy*/

func canConstruct(ransomNote string, magazine string) bool {
	record := make([]int, 26)
	for _, v := range magazine {
		record[v-'a']++
	}
	for _, v := range ransomNote {
		record[v-'a']--
		// magazine 中的每个字符只能在 ransomNote 中使用一次
		if record[v-'a'] < 0 {
			return false
		}
	}
	return true
}

// func canConstruct(ransomNote string, magazine string) bool {
// 	m := make(map[int]int)
//
// 	for _, v := range magazine {
// 		m[int(v-'a')]++
// 	}
//
// 	for _, v := range ransomNote {
// 		m[int(v-'a')]--
// 		if _, ok := m[int(v-'a')]; !ok {
// 			return false
// 		}
// 		m[int(v-'a')]--
// 		if m[int(v-'a')] < 0 {
// 			return false
//
// 		}
// 	}
//
// 	return true
//
// }
