package main

import "fmt"

func main() {
	// 普通for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 范围遍历
	arr := [3]int{100, 200, 300}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// 范围map
	// 创建一个长度的map    map[keyType]valueType
	nameMap := make(map[string]int, 1)
	// 赋值
	nameMap["小米"] = 10
	// 赋值，会自动扩容
	nameMap["小红"] = 18
	for k, v := range nameMap {
		fmt.Println(k, v)
	}

	// while
	num := 0
	for num < 10 {
		fmt.Println(num)
		num++
	}

	// 死循环
	//for {
	//	fmt.Println("go go study, day day up")
	//}
}
