package main

//
// /*232-用栈实现队列*/
//
// type MyQueue struct {
// 	arr []int
// }
//
// func Constructor() MyQueue {
// 	return MyQueue{[]int{}}
// }
//
// func (this *MyQueue) Push(x int) {
// 	this.arr = append(this.arr, x)
// }
//
// func (this *MyQueue) Pop() int {
// 	x := this.arr[0]
// 	this.arr = this.arr[1:]
// 	return x
// }
//
// func (this *MyQueue) Peek() int {
// 	return this.arr[0]
// }
//
// func (this *MyQueue) Empty() bool {
// 	if len(this.arr) == 0 {
// 		return true
// 	}
// 	return false
// }
//
// /**
//  * Your MyQueue object will be instantiated and called as such:
//  * obj := Constructor();
//  * obj.Push(x);
//  * param_2 := obj.Pop();
//  * param_3 := obj.Peek();
//  * param_4 := obj.Empty();
//  */
