package main

// // 单行注释
//#include <stdio.h>
//void say1(){
// printf("Hello Word1\n");
//}
// // 不能有空行，但可以有注释
/*
// 多行注释
#include <stdio.h>
void say2(){
 printf("Hello Word2\n");
}
*/
import "C"

// 如果想在 go 中使用 c 语言，要借助 C 包 和注释
func main() {
	C.say1()
	C.say2()

	// 注意：
	// 1、单行注释要这样写 // // 单行注释
	// 2、import "C"  和 c 代码之间不能有空行，但可以有注释
}
