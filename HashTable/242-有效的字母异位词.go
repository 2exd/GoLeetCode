package main

/*242. 有效的字母异位词 easy*/
// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	exists := make(map[byte]int)
// 	for i := 0; i < len(s); i++ {
// 		if v, ok := exists[s[i]]; v >= 0 && ok {
// 			exists[s[i]] = v + 1
// 		} else {
// 			exists[s[i]] = 1
// 		}
// 	}
// 	for i := 0; i < len(t); i++ {
// 		if v, ok := exists[t[i]]; v >= 1 && ok {
// 			exists[t[i]] = v - 1
// 		} else {
// 			return false
// 		}
// 	}
// 	return true
// }

func isAnagram(s string, t string) bool {
	lenS, lenT := len(s), len(t)
	if lenS != lenT {
		return false
	}
	ms := make(map[int32]int, lenS)
	for _, c := range s {
		ms[c]++
	}
	for _, v := range t {
		if _, ok := ms[v]; !ok {
			return false
		}
		ms[v]--
		if ms[v] < 0 {
			return false
		}
	}
	return true
}
