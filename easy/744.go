package main

/*744. 寻找比目标字母大的最小字母*/
/* 如果letter大于target，就直接找比letter大的
如果letter小于target，*/
func nextGreatestLetter(letters []byte, target byte) byte {
	var minDistance = 199
	for _, letter := range letters {
		if letter > target {
			minDistance = findMinDistance(minDistance, int(letter-target))
		} else if letter < target {
			minDistance = findMinDistance(minDistance, int(letter+26-target))
		}
	}
	ret := (minDistance+int(target)-96)%26 + 96
	if ret == 96 {
		return byte(ret + 26)
	}
	return byte(ret)
}

func findMinDistance(x, y int) int {
	if x > y {
		return y
	} else {
		return x

	}
}

// func main() {
//	//var letters = []byte{'c', 'f', 'j'}
//	//target := 'a'
//	//var letters = []byte{'c', 'f', 'j'}
//	//target := 'j'
//
//	var letters = []byte{'a','y','z'}
//	target := 'y'
//	fmt.Printf("%d\n", 'z')
//	fmt.Printf("%d\n", '}')
//	fmt.Printf("%d\n", '`')
//	fmt.Printf("%c", nextGreatestLetter(letters, byte(target)))
// }
