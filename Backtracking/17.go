package main

var (
	m      []string
	path17 []byte
	res17  []string
)

func letterCombinations(digits string) []string {
	m = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path17, res17 = make([]byte, 0), make([]string, 0)
	if digits == "" {
		return res17
	}
	dfs17(digits, 0)
	return res17
}

func dfs17(digtis string, k int) {
	if len(path17) == len(digtis) {
		tmp := string(path17)
		res17 = append(res17, tmp)
		return
	}

	for i := k; i < len(digtis); i++ {
		num := digtis[i] - '0'
		str := m[num-2]
		for j := 0; j < len(str); j++ {
			path17 = append(path17, str[j])
			dfs17(digtis, i+1)
			path17 = path17[:len(path17)-1]
		}
	}

}
