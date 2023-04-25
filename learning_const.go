package main

import "fmt"

func main() {
	const (
		num1 = 998
		num2 // 和上一行的值一样
		num3 = 666
		num4 // 和上一行的值一样
		num5 // 和上一行的值一样
	)
	fmt.Println("num1 = ", num1) // 998
	fmt.Println("num2 = ", num2) // 998
	fmt.Println("num3 = ", num3) // 666
	fmt.Println("num4 = ", num4) // 666
	fmt.Println("num5 = ", num5) // 666
	const (
		str1, str2 = "aaa", "bbb"
		str3, str4 //和上一行的值一样，注意变量个数必须也和上一行一样
	)
	fmt.Println("str1 =", str1) // aaa
	fmt.Println("str2 =", str2) // bbb
	fmt.Println("str3 =", str3) // aaa
	fmt.Println("str4 =", str4) // bbb
}
