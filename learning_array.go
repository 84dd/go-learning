package main

import (
	"fmt"
)

func main() {
	learningArr1()
	learningArr2()
	learningSlice()
}

// 一维数组的定义
func learningArr1() {
	// 定义数组格式： var 变量名 [数组长度]数据类型
	var arr1 [3]int   // 默认填充0
	fmt.Println(arr1) // 0 0 0
	arr2 := [3]int{}  // 默认填充0
	fmt.Println(arr2) // 0 0 0

	// 定义并初始化
	arr3 := [3]int{1, 3, 5}
	fmt.Println(arr3) // 1 3
	arr4 := [3]int{1, 3}
	fmt.Println(arr4) // 1 3 0
	arr5 := [3]int{0: 1, 2: 5}
	fmt.Println(arr5) // 1 0 5
	arr6 := [3]int{}
	arr6[1] = 3
	fmt.Println(arr6) // 0 3 0

	// 不定长度
	arr7 := [...]int{1, 3, 5}
	fmt.Println(arr7) // 1 3 5
	arr8 := [...]int{0: 1, 2: 5}
	fmt.Println(arr8) // 1 0 5

	// 遍历
	arr9 := [...]int{1, 3, 5}
	for i := 0; i < len(arr9); i++ {
		fmt.Println(arr9[i])
	}
	for i, v := range arr9 {
		fmt.Println(i, v)
	}

	// 元素比较（同类型的数组才能比较，否则编译报错）
	arr10 := [3]int{1, 2, 3}
	arr11 := [3]int{1, 2, 4}
	arr12 := [3]int{1, 2, 3}
	fmt.Println(arr10 == arr11) // false
	fmt.Println(arr10 == arr12) // true

	// Go语言中的数组是值类型,赋值和传参时会复制整个数组
	arr13 := [3]int{1, 2, 3}
	arr14 := arr13
	arr14[2] = 666
	fmt.Println(arr13) // 1 2 3
	fmt.Println(arr14) // 1 2 666

}

// 二维数组
func learningArr2() {
	arr1 := [2][3]int{}
	fmt.Println(arr1) // [[0 0 0] [0 0 0]]

	// ... 只能在第一维使用
	arr2 := [...][3]int{
		{1, 1, 1},
		{},
		{1, 5}, // 因为换行，最后一个逗号需保留
	}
	fmt.Println(arr2) // [[1 1 1] [0 0 0] [1 5 0]]
}

// 切片，同python
// 切片源码如下
//
//	type slice struct {
//		array unsafe.Pointer // 指向底层数组指针
//		len   int            // 切片长度(保存了多少个元素)
//		cap   int            // 切片容量(可以保存多少个元素)
//	}
func learningSlice() {
	// 切片，对原数组不影响
	arr1 := [5]int{1, 3, 5, 7, 9}

	// [开始位置:结束位置) -> 注意开闭情况，这里相当于取 0 1 下标
	arr2 := arr1[0:2]
	// [开始位置:结束位置) -> 注意开闭情况，这里相当于取 1 2 下标
	arr3 := arr1[1:3]
	fmt.Println(arr1) // 1 3 5 7 9
	fmt.Println(arr2) // 1 3
	fmt.Println(arr3) // 3 5

	arr4 := arr1[1:]
	arr5 := arr1[:3]
	arr6 := arr1[:]
	fmt.Println(arr4) // 3 5 7 9
	fmt.Println(arr5) // 1 3 5
	fmt.Println(arr6) // 1 3 5 7 9

	// 由于切片指向的是数组指针，所以如果修改原数组的值，切片也会跟着变化
	fmt.Println(arr1) // 1 3 5 7 9
	// arr2 := arr1[0:2]
	fmt.Println(arr2) // 1 3
	arr1[0] = 666
	fmt.Println(arr2) // 666 3

	// append：切片最佳数据
	// copy：赋值切片
}
