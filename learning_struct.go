package main

import (
	"fmt"
	"go-learning/src/code.lanym.com/entity"
)

// struct：结构体，即java中的对象
func main() {
	learnStruct1()
	learnStruct2()
}

// 普通用法
func learnStruct1() {
	var student1 entity.Student
	student1.Name = "小明"
	student1.Age = 18
	fmt.Println(student1)
	// 普通对象
	var student2 = entity.Student{Name: "小明", Age: 19}
	student1.Age = 20
	fmt.Println(student2)
}

// 匿名用法
func learnStruct2() {
	// 定义临时结构体
	var student3 = struct {
		name string
		age  int
		sex  int
	}{
		name: "小英",
	}
	fmt.Println(student3)

	// 复杂结构体
	var project1 entity.Project
	fmt.Println(project1)
	project1.Student.Name = "小明"
	project1.Stu2.Name = "小红"
	project1.Status = 4
	project1.Remark = "备注"
	fmt.Println(project1)
}
