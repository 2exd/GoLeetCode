package main

import (
"fmt"
"math"
)

func main() {

	fmt.Printf("float64的最大值是:%.f\n", math.MaxFloat64)
	fmt.Printf("float64的最小值是:%.f\n", math.SmallestNonzeroFloat64)
	fmt.Printf("float32的最大值是:%.f\n", math.MaxFloat32)
	fmt.Printf("float32的最小值是:%.f\n", math.SmallestNonzeroFloat32)
	fmt.Printf("Int8的最大值是:%d\n", math.MaxInt8)
	fmt.Printf("Int8的最小值是:%d\n", math.MinInt8)
	fmt.Printf("Uint8的最大值是:%d\n", math.MaxUint8)
	fmt.Printf("Int16的最大值是:%d\n", math.MaxInt16)
	fmt.Printf("Int16的最小值是:%d\n", math.MinInt16)
	fmt.Printf("Uint16的最大值是:%d\n", math.MaxUint16)
	fmt.Printf("Int32的最大值是:%d\n", math.MaxInt32)
	fmt.Printf("Int32的最小值是:%d\n", math.MinInt32)
	fmt.Printf("Uint32的最大值是:%d\n", math.MaxUint32)
	fmt.Printf("Int64的最大值是:%d\n", math.MaxInt64)
	fmt.Printf("Int64的最小值是:%d\n", math.MinInt64)
	//fmt.Println("Uint64的最大值是:", math.MaxUint64)
	fmt.Printf("圆周率默认为:%.200f\n", math.Pi)

}