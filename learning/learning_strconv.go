package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 第一个参数：要转换的内容
	// 第二个参数：转换到什么进制
	// 第三个参数：转换为多少位整形
	// 返回结果： 转换结果，报错信息
	fmt.Println(strconv.ParseInt("123", 10, 8))
	// 报错输出 0 strconv.ParseInt: parsing "aaa": invalid syntax
	fmt.Println(strconv.ParseInt("aaa", 10, 8))
	//strconv.ParseInt
	//strconv.ParseBool
	//strconv.ParseFloat
	//strconv.ParseComplex
	//strconv.ParseUint

	// 第一个参数：要转换的内容
	// 第二个参数：转换到什么进制
	// 返回结果： 转换结果
	fmt.Println(strconv.FormatInt(123, 2))
	//strconv.FormatInt
	//strconv.FormatBool
	//strconv.FormatFloat
	//strconv.FormatComplex
	//strconv.FormatUint

	// 便捷写法
	fmt.Println(strconv.Itoa(123))
	fmt.Println(strconv.Atoi("123"))
	fmt.Println(strconv.Atoi("123aaa"))

	// 利用 Sprintf: Sprintf和Println很像，但Sprintf是格式化并返回，Println是格式化并输出
	fmt.Println(fmt.Sprintf("%d", 123))
	fmt.Println(fmt.Sprintf("%f", 123.1))

}
