package main

import (
	"fmt"
)

func main() {
	// 定义切片泛型
	type mySlice[T int | string] []T
	var ms = mySlice[int]{1, 2, 3}
	fmt.Printf("切片泛型:%v\n", ms)
	//定义map
	type myMap[Key string | int, Value string | int] map[Key]Value
	var mm = myMap[string, int]{"math": 90, "chinese": 80}
	fmt.Printf("map泛型：%v\n", mm)
	//定义机构体
	type myStruct[T string | int, TT int | float64] struct {
		Name    string
		Content T
		Wallet  TT
	}
	var mst = myStruct[string, int]{Name: "hll", Content: "无内容", Wallet: 100}
	fmt.Printf("struct泛型：%v\n", mst)

	// 定义泛型函数
	fmt.Printf("调用泛型函数：%d\n", sum(1, 4))
	fmt.Printf("调用泛型函数：%f\n", sum(1.2, 4.8))

	//自定义类型约束
	// 并集约束
	forEach([]int{1, 2, 3})
	forEach([]float64{1.1, 2.2, 3.3})

	// 约束交集
	forEach1([]int8{1, 2, 3})
	//forEach1([]int{1, 2, 3}) //报错

	//类型的超集 ~： 约束的范围也包含只要最底层的类型
	fmt.Printf("调用泛型超集函数：%d\n", superSum(1, 4))

	fmt.Printf("调用泛型超集函数：%d\n", superSum(def_int(1), def_int(4)))

}

// 定义泛型函数
func sum[T int | float64](a, b T) T {
	return a + b
}

//自定义类型约束
type myInt interface {
	int | int8 | int16 | int32 | int64
}
type myUInt interface {
	uint | uint8 | uint16 | uint32 | int8
}
type myFloat interface {
	float32 | float64
}

// 1. 约束并集   约束范围是：myInt + myFloat
type myNumber interface {
	myInt | myFloat
}

// 交集约束 约束范围是：myInt - myFloat
type myNumber1 interface {
	myInt
	myUInt
}

// 遍历切片 约束并集
func forEach[T myNumber](l []T) {
	for _, t := range l {
		fmt.Println(t)
	}
}

// 遍历切片 约束并集
func forEach1[T myNumber1](l []T) {
	for _, t := range l {
		fmt.Println(t)
	}
}

// 超集
type mySuperInt interface {
	//int // 这样定义是不允许 def_int 作为参数的
	~int
}
type def_int int

func superSum[T mySuperInt](a, b T) T {
	return a + b
}
