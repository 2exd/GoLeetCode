package main

type MinStack struct {
	arr    []int
	minArr []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.arr = append(this.arr, val)
	n := len(this.minArr)
	// len is 0 or last element bigger than val
	if n == 0 || this.minArr[n-1] > val {
		// append to stack end
		this.minArr = append(this.minArr, val)
	} else {
		// append min to stack end
		this.minArr = append(this.minArr, this.minArr[n-1])
	}
}

func (this *MinStack) Pop() {
	// pop the last element of stacks
	this.arr = this.arr[:len(this.arr)-1]
	this.minArr = this.minArr[:len(this.minArr)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	return this.minArr[len(this.minArr)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
