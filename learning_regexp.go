package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 提取数字
	str1 := "aaa123你我他456"
	compile, _ := regexp.Compile("\\d+")
	res := compile.FindAllString(str1, -1)
	fmt.Println(res) //[123 456]

	fmt.Println(compile.Match([]byte(str1)))     // true  因为包含了数字
	fmt.Println(compile.Match([]byte("123321"))) // true

	compile2, _ := regexp.Compile("^\\d+$")
	fmt.Println(compile2.Match([]byte(str1))) // false  因为不是全部数字

}
