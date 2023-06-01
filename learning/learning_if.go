package main

import "fmt"

func main() {
	// 情况1：可以在 if 同一行写变量的定义
	if age := 18; age < 20 {
		fmt.Println(age)
	} else {
		fmt.Println(age)
	}

	// 情况2：可以调用函数
	if age := getAge(); age < 20 {
		fmt.Println(age)
	} else {
		fmt.Println(age)
	}
}

func getAge() int {
	return 18
}
