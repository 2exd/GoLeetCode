package main

//
// import (
// 	"fmt"
// 	Iterator "zexd.com/goLeetCode/DataStructure/utils/iterator"
// )
//
// func main() {
// 	var arr = make([]interface{}, 0, 0)
// 	arr = append(arr, 5)
// 	arr = append(arr, 3)
// 	arr = append(arr, 2)
// 	arr = append(arr, 4)
// 	arr = append(arr, 1)
// 	arr = append(arr, 4)
// 	arr = append(arr, 3)
// 	arr = append(arr, 1)
// 	arr = append(arr, 5)
// 	arr = append(arr, 2)
//
// 	// new iterator
// 	i := Iterator.New(&arr)
//
// 	// from the beginning
// 	fmt.Println("begin")
// 	for i := i.Begin(); i.HasNext(); i.Next() {
// 		fmt.Println(i.Value())
// 	}
// 	fmt.Println()
//
// 	// from the end
// 	fmt.Println("end")
// 	for i := i.End(); i.HasPre(); i.Pre() {
// 		fmt.Println(i.Value())
// 	}
// 	fmt.Println()
//
// 	// from 4
// 	fmt.Println("get4")
// 	for i := i.Get(4); i.HasNext(); i.Next() {
// 		fmt.Println(i.Value())
// 	}
// 	fmt.Println()
// }
