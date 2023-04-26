package main

import (
	"fmt"
	"go-learning/src/code.lanym.com/other"
	"regexp"
	"strconv"
)

func main() {
	// 调用方法
	var age = 18
	myFun("小明", age)
	addAge("小明", &age)
	fmt.Println("调用函数 addAge ，会改变地址的值，age=", age)

	i2, i3, _, i5 := calc1(1, 2)
	fmt.Println(i2, i3, i5)

	add, _, multiply, divide := calc2(2, 3)
	fmt.Println(add, multiply, divide)

	// 调用其他包的方法
	other.FinelyPrint("小红", 3)

	// 匿名函数
	func() {
		fmt.Println("hello word")
	}()
	word := func(age int) string {
		return "你今年" + strconv.Itoa(age) + "岁了"
	}(18)
	fmt.Println(word)

	// 在 go 中，func可以作为参数，也可以作为返回值，也可以创建
	// 类似js
	theAgeFun := func(ageStr string) func() int64 {
		sampleRegexp := regexp.MustCompile(`\d+`)
		if sampleRegexp.MatchString(ageStr) {
			return func() int64 {
				parseInt, _ := strconv.ParseInt(ageStr, 10, 4)
				return parseInt
			}
		} else {
			return func() int64 {
				return 0
			}
		}
	}
	numAge := theAgeFun("12")
	fmt.Println("转换到的数字年龄为：", numAge())

}

// myFun 定义方法
//
// param: name
// param: age
func myFun(name string, age int) {
	fmt.Printf("%v今年%v岁\n", name, age)
}

// addAge 传递指针
//
// param: name
// param: *age
func addAge(name string, age *int) {
	*age++
	fmt.Printf("%v明年%v岁\n", name, *age)
}

// calc 多返回方法示例
// 方法后面定义返回结果的数量和类型
// param: num1
// param: num2
func calc1(num1 int, num2 int) (int, int, int, int) {
	// 返回的结果可以用var声明
	var add = num1 + num2
	// 也可以是临时变量
	subtract := num1 - num2
	var multiply = num1 * num2
	var divide = num1 / num2
	return add, subtract, multiply, divide
}

// calc 多返回方法示例
// 方法后面也可以指定返回结果的名字
// 这种定义方式，在开发工具中，调用函数按 tab 键会自动补全输出结果的变量名
// param: num1
// param: num2
func calc2(num1 int, num2 int) (add int, subtract int, multiply int, divide int) {
	add = num1 + num2
	subtract = num1 - num2
	multiply = num1 * num2
	divide = num1 / num2
	return add, subtract, multiply, divide
}
