package main

// 
// import (
// 	"fmt"
// 	"sync"
// 	"zexd.com/goLeetCode/DataStructure/list"
// )
//
// func main() {
// 	l := list.New()
// 	wg := sync.WaitGroup{}
// 	// initialize
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func(idx uint64) {
// 			l.Insert(idx, idx)
// 			wg.Done()
// 		}(uint64(i))
// 	}
// 	wg.Wait()
//
// 	fmt.Println("使用迭代器访问全部元素；")
// 	for i := l.Iterator(); i.HasNext(); i.Next() {
// 		fmt.Println(i.Value())
// 	}
// 	l.Set(5, "测试")
// 	fmt.Println("输出刚设定的测试元素的位置：", l.IndexOf("测试"))
// 	fmt.Println("从测试位生产新list，长度上限为10，并从头部输出：")
// 	newList := l.SubList(l.IndexOf("测试"), 10)
// 	length := newList.Size()
// 	fmt.Println("新链表的长度：", length)
// 	for i := uint64(0); i < length; i++ {
// 		fmt.Println(newList.Get(0))
// 		newList.Erase(0)
// 	}
// 	fmt.Println("从结尾向首部输出原链表：")
// 	for i := l.Size(); i > 0; i-- {
// 		fmt.Println(l.Get(i - 1))
// 		l.Erase(i - 1)
// 	}
// }
