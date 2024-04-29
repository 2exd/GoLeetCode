package main

type MemoryBlock struct {
	Start, End int
}

func solu() {
	// sc := bufio.NewReader(os.Stdin)
	// input, _ := sc.ReadString('\n')
	// n, _ := strconv.Atoi(strings.TrimSpace(input))
	// var mb []MemoryBlock
	//
	// for {
	// 	input, err := sc.ReadString('\n')
	// 	if err != nil || strings.TrimSpace(input) == "" {
	// 		break
	// 	}
	// 	parts := strings.Fields(input)
	// 	offset, _ := strconv.Atoi(parts[0])
	// 	size, _ := strconv.Atoi(parts[1])
	// 	mb = append(mb, MemoryBlock{Start: offset, End: offset + size - 1})
	// }
	//
	// sort.Slice(mb, func(i, j int) bool {
	// 	return mb[i].Start < mb[j].Start
	// })
	//
	// isValid := true
	// for i := 0; i < len(mb)-1; i++ {
	// 	start := mb[i].Start
	// 	end := mb[i].End
	// 	if start < 0 || end >= 100 || start > end {
	// 		isValid = false
	// 		break
	// 	}
	// 	if i > 0 && mb[i-1].End >= start {
	// 		isValid = false
	// 		break
	// 	}
	// }
	//
	// if !isValid {
	// 	fmt.Println(-1)
	// 	return
	// }
	//
	// offset := -1
	// maxDistance := 101
	//
	// if mb[0].Start > n {
	// 	offset = 0
	// 	maxDistance = mb[0].Start - n
	// }

}

// func main() {
// 	solu()
// }
