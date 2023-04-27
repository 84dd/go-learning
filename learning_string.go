package main

import (
	"fmt"
	"strings"
)

func main() {
	// go 语言是 utf-8编码，一个中文占3字节长度
	fmt.Println(len("abc")) // 3
	fmt.Println(len("你我他")) // 9

	// byte占一个字节，而中文占3个字节，所以byte不能保存中文
	arr1 := []byte("你我他")
	fmt.Println(len(arr1)) // 9
	for _, v := range arr1 {
		fmt.Printf("%c", v) // 乱码
	}
	fmt.Println()

	// rune类型就是用来专门保存中文的
	arr2 := []rune("你我他")
	fmt.Println(len(arr2)) // 3
	for _, v := range arr2 {
		fmt.Printf("%c", v) // 正常
	}
	fmt.Println()

	// 字符串的相关操作在 strings包中
	// 具体用法和js或java大同小异
	index := strings.Index("你我他", "狗")
	fmt.Println(index)

	str1 := "你我他"
	str2 := "你我他"
	fmt.Println("值比较 ", str1 == str2)   // true
	fmt.Println("地址比较", &str1 == &str2) // false

}
