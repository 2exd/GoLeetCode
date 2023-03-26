package main

/*考察元素入栈出栈顺序是否合理问题*/
// func main() {
//
// }

func isValidStackSequence(in, out []int) bool {
	stack := make([]int, len(in))
	inLen, outLen := len(in), len(out)
	if inLen != outLen {
		return false
	}
	inIdx, outIdx := 0, 0
	for outIdx < outLen {
		if len(stack) > 0 && stack[len(stack)-1] == out[outIdx] { // 如果栈顶元素和出站火车一致
			stack = stack[:len(stack)-1] // 将栈顶元素弹出
			outIdx++                     // 出站火车下标加1
		} else if inIdx < inLen { // 如果栈顶元素和出站火车不一致，且还有火车可以入站
			stack = append(stack, in[inIdx]) // 将入站火车入栈
			inIdx++                          // 入站火车下标加1
		} else { // 如果栈顶元素和出站火车不一致，且没有火车可以入站
			return false // 调度失败，返回 false
		}
	}
	return true
}

//
// func main() {
// 	inputSequence := []int{1, 2, 3, 4, 5}
// 	outputSequence := []int{4, 5, 3, 2, 1}
// 	isValid := isValidStackSequence(inputSequence, outputSequence)
// 	if isValid {
// 		fmt.Print("Yes")
// 	} else {
// 		fmt.Print("No")
// 	}
// }
