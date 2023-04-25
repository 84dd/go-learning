package main

import (
	"fmt"
	"net/http"
	"unsafe"
)

// 定义全局变量
var aaa = 10

// 定义全局不可以修改的变量
const bbb = 20

func main() {
	aaa = 20
	fmt.Println(aaa)
	// bbb = 30
	fmt.Println(bbb)

	fmt.Println("int 占用的空间：", unsafe.Sizeof(int(0)))
	fmt.Println("int64 占用的空间：", unsafe.Sizeof(int64(0)))
	fmt.Println("boolean 占用的空间：", unsafe.Sizeof(true))

	// 声明，默认会赋零值
	var a int
	fmt.Println(a) // 0

	var b string
	fmt.Println(b) // 空字符串

	var c *int
	fmt.Println(c) // nil

	var d, e, f int      // int, int, int
	fmt.Println(d, e, f) // 0，每一个都是int类型，并且赋0值

	var h, i, j = true, 2.3, "four" // bool, float64, string
	fmt.Println(h, i, j)            // 同时赋值，并自动类型推断

	k := 0.0 // 临时变量的简短声明，不用带 var
	fmt.Println(k)

	l, m := 0, "i am string" // 同时赋多个临时变量，并自动类型推断
	fmt.Println(l, m)

	p, q := 111, 222
	p, q = q, p // 交換值
	fmt.Println(p, q)

	url := "http://wx.qlogo.cn/Vaz7vE1/64"
	resp1, err1 := http.Get(url) // 返回两个值，用两个变量分别接收
	fmt.Println(resp1, err1)
	_, err2 := http.Get(url) // 返回两个值，但第一个值用 _ 下划线，表示丢弃不需要的值
	fmt.Println(err2)

	// 变量组
	var (
		num1 int
		str1 string
	)
	num1 = 10
	str1 = "aaa"
	fmt.Println(num1, str1)

	// 变量组
	var (
		num2        int = 10
		str2, bool2     = "", true
	)
	fmt.Println(num2, str2, bool2)

}
