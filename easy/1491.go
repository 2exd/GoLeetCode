package main

/*1491. 去掉最低工资和最高工资后的工资平均值*/

func average(salary []int) float64 {
	var computeAver func([]int) float64
	computeAver = func(arr []int) float64 {
		var sum int
		for i := 1; i < len(arr)-1; i++ {
			sum += arr[i]
		}
		return float64(sum) / float64((len(arr) - 2))
	}

	// 堆排序
	//var HeapSort func([]int) []int
	//var HeapSortMax func ([]int, int) []int
	//HeapSortMax = func(arr []int, length int) []int {
	//	// length := len(arr)
	//	if length <= 1 {
	//		return arr
	//	}
	//	depth := length/2 - 1 //二叉树深度
	//	for i := depth; i >= 0; i-- {
	//		topmax := i //假定最大的位置就在i的位置
	//		leftchild := 2*i + 1
	//		rightchild := 2*i + 2
	//		if leftchild <= length-1 && arr[leftchild] > arr[topmax] { //防止越过界限
	//			topmax = leftchild
	//		}
	//		if rightchild <= length-1 && arr[rightchild] > arr[topmax] { //防止越过界限
	//			topmax = rightchild
	//		}
	//		if topmax != i {
	//			arr[i], arr[topmax] = arr[topmax], arr[i]
	//		}
	//	}
	//	return arr
	//}
	//
	//HeapSort = func(arr []int) []int {
	//	length := len(arr)
	//	for i := 0; i < length; i++ {
	//		lastlen := length - i
	//		HeapSortMax(arr, lastlen)
	//		if i < length {
	//			// 交换首尾节点(为了维持一个完全二叉树才要进行收尾交换)
	//			arr[0], arr[lastlen-1] = arr[lastlen-1], arr[0]
	//		}
	//	}
	//	return arr
	//}
	//return	computeAver(HeapSort(salary))

	// 快速排序
	//var QuickSort func(arr []int) []int
	//QuickSort = func(arr []int) []int {
	//	if len(arr) <= 1 {
	//	return arr
	//}
	//	splitdata := arr[0]          //第一个数据
	//	low := make([]int, 0, 0)     //比我小的数据
	//	hight := make([]int, 0, 0)   //比我大的数据
	//	mid := make([]int, 0, 0)     //与我一样大的数据
	//	mid = append(mid, splitdata) //加入一个
	//
	//	for i := 1; i < len(arr); i++ {
	//	if arr[i] < splitdata {
	//	// 加入low
	//	low = append(low, arr[i])
	//} else if arr[i] > splitdata {
	//	// 加入height
	//	hight = append(hight, arr[i])
	//} else {
	//	// 加入mid
	//	mid = append(mid, arr[i])
	//}
	//}
	//	low, hight = QuickSort(low), quickSort(hight)
	//	myarr := append(append(low, mid...), hight...)
	//	return myarr
	//}
	//return	computeAver(QuickSort(salary))

	// 冒泡排序
	//var BubbleSort func([]int) []int
	//BubbleSort = func(arr []int) []int {
	//	for i := 0; i < len(arr); i++ {
	//		for j := i + 1; j < len(arr); j++ {
	//			if arr[i] > arr[j] {
	//				arr[i], arr[j] = arr[j], arr[i]
	//			}
	//		}
	//	}
	//	return arr
	//}
	//return	computeAver(BubbleSort(salary))

	// 基数排序

	// 选择排序
	//var SelectSort func ([]int) []int
	//SelectSort = func (arr []int) []int{
	//	length := len(arr)
	//	fmt.Println(length)
	//	if length <= 1 {
	//		return arr
	//	}
	//	for i := 0; i < length; i++ {
	//		min := i
	//		for j := i + 1; j < length; j++ {
	//			if arr[min] > arr[j] {
	//				min = j
	//			}
	//		}
	//		if i != min {
	//			arr[i], arr[min] = arr[min], arr[i]
	//		}
	//	}
	//	return arr
	//}
	//selectsort := SelectSort(salary)
	//fmt.Println(SelectSort(selectsort))
	//
	//return	computeAver(SelectSort(salary))

	// 基数排序
	//var RadixSort func (arr []int) []int
	//var CountingSrot func(arr []int, loc int)
	//var digit func(num int, loc int) int
	//RadixSort = func (arr []int) []int {
	//	maxValueLen := 0 // 需要循环多少次，以最大数字为准
	//	for i := 0; i < len(arr); i++ {
	//	n := len(strconv.Itoa(arr[i])) // 方便起见，数字转字符，再取长度
	//	if n > maxValueLen {
	//	maxValueLen = n
	//}
	//}
	//	for loc := 1; loc <= maxValueLen; loc++ {
	//		CountingSrot(arr, loc)
	//}
	//	return arr
	//}
	//
	//CountingSrot = func (arr []int, loc int) {
	//	bucket := make([][]int, 10) // 0~9 总共10个队列
	//
	//	for i := 0; i < len(arr); i++ {
	//		ji := digit(arr[i], loc) // 获取对应位的数字
	//		if bucket[ji] == nil {
	//			bucket[ji] = make([]int, 0) // 如果 bucket 需要再去初始化
	//		}
	//
	//		bucket[ji] = append(bucket[ji], arr[i]) // 将数字 push 进去
	//	}
	//
	//	idx := 0
	//	for i := 0; i <= 9; i++ {
	//		for j := 0; j < len(bucket[i]); j++ {
	//			// 遍历二维数组
	//			arr[idx] = bucket[i][j] //将数字取出来 给原数组重新赋值
	//			idx++
	//		}
	//	}
	//}
	//
	//// 数字，右数第几位，从1开始
	//digit = func(num int, loc int) int {
	//	return num % int(math.Pow10(loc)) / int(math.Pow10(loc-1))
	//}
	//
	//return	computeAver(RadixSort(salary))

	// 归并排序
	//var merge = func (arr []int, start, mid, end int) {
	//	var tmparr = []int{}
	//	var s1, s2 = start, mid+1
	//	for s1<= mid && s2<= end{
	//		if arr[s1] > arr[s2] {
	//			tmparr = append(tmparr, arr[s2])
	//			s2++
	//		} else {
	//			tmparr = append(tmparr, arr[s1])
	//			s1++
	//		}
	//	}
	//	if s1<=mid {
	//		tmparr = append(tmparr, arr[s1: mid+1]...)
	//	}
	//	if s2<=end {
	//		tmparr = append(tmparr, arr[s2: end+1]...)
	//	}
	//	for pos,item:=range tmparr{
	//		arr[start + pos] = item
	//	}
	//}
	//var MergeSort func (arr []int, start, end int)
	//MergeSort = func (arr []int, start, end int){
	//	if start >= end {
	//		return
	//	}
	//	mid:=(start + end) / 2
	//	MergeSort(arr, start, mid)
	//	MergeSort(arr, mid+1, end)
	//	merge(arr, start, mid, end)
	//}
	//MergeSort(salary, 0, len(salary)-1)
	//return computeAver(salary)

	// 插入排序
	//var InsertSort func(st []int) []int
	//InsertSort = func(st []int) []int {
	//	if len(st) <= 1 {
	//	return st
	//}
	//	for i := 1; i < len(st); i++ {
	//	back := st[i]
	//	j := i - 1
	//	for j >= 0 && back < st[j] {
	//	st[j+1] = st[j]
	//	j--
	//}
	//	st[j+1] = back
	//}
	//	return st
	//}
	//return computeAver(InsertSort(salary))

	// 希尔排序
	var ShellSort = func(num []int) {

		//increment相隔数量
		for increment := len(num) / 2; increment > 0; increment /= 2 {
			//i序号较大的数组下标，i ,j进行比较

			for i := increment; i < len(num); i++ {
				//进行交换
				temp := num[i]
				//按照increment，数组从j到0进行交换比较
				for j := i - increment; j >= 0; j -= increment {
					if temp < num[j] {
						num[j+increment] = num[j]
						num[j] = temp
						temp = num[j]
					} else { //由于数组前面按照increment已经排好序，如果temp>num[j],则不必继续比较交换下去
						break
					}
				}

			}

		}
	}
	ShellSort(salary)
	return computeAver(salary)
}

//func main()  {
//	//arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
//	//arr := []int{48000,59000,99000,13000,78000,45000,31000,17000,39000,37000,93000,77000,33000,28000,4000,54000,67000,6000,1000,11000}
//	arr := []int{4000,3000,1000,2000}
//	fmt.Println(average(arr))
//}
