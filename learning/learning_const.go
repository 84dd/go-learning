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

	const (
		age1 = iota
		age2 = iota
		age3 = 10   // 可以中断 iota，赋新的值
		age4        // 如果不恢复 iota，则这一行使用上一行的值 10
		age5 = iota // 恢复 iota 后，递增到第4行，所以赋值4
	)
	fmt.Println(age1, age2, age3, age4, age5) // 0 1 10 10 4

	const (
		page1, page2 = iota, iota // 可以多个赋值，iota
		page3, page4 = iota, 10   // 也可以一个iota，另一个赋具体指
		page5, page6              // 都不指定值，则第一个变量使用上一行的iota，第二个变量使用 上一个的第二个值
		page7, page8 = iota, iota // iota 和行数有关
	)

}
