package lesson

import (
	"fmt"
	"strconv"
)

// 学习 go 语言的函数，变量，函数调用等基本用法，然后完成如下：
// 参考章节 5-1函数声明
// 前提： a 和 b 都是 int32
// 在之前提交作业的 main.go 文件内 (或者新建 lesson2.go 文件)
// 1.	新建方法add：入参为 a 和 b，返回两者的和
// 2.	新建方法sub：入参为 a 和 b，返回两者的差
// 3.	新建方法mult：入参为 a 和 b，返回两者的乘积
// 4.	新建方法div：入参为 a 和 b，返回两者的商，保留两位小数

func add(a int32, b int32) int32 {
	return a + b
}

func sub(a int32, b int32) int32 {
	return a - b
}

func mult(a int32, b int32) int32 {
	return a * b
}

func div(a int32, b int32) float64 {
	// 考察点1： int 类型的数如果直接相处，结果也是int类型，和java类似，所以要 float64(a)/float64(b)
	// 考察点2： 保留两位小时用 fmt.Sprintf（字符串）
	// 考察点3： 字符串转数字用 strconv.ParseFloat
	num, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(a)/float64(b)), 64)
	return num
}
