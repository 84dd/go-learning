package main

import (
	"fmt"
	"go-learning/src/code.lanym.com/entity"
)

// struct：结构体，即java中的对象
func main() {
	learnStruct1()
	learnStruct2()
	learnStruct3()
}

// 普通用法
func learnStruct1() {
	var student1 entity.Student
	student1.Name = "小明"
	student1.Age = 18
	fmt.Println(student1)
	// 普通对象
	var student2 = entity.Student{Score: 10}
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

// 封装对象
func learnStruct3() {
	person1 := entity.Person{Name: "小明", Age: 18}
	student := entity.Student{Person: person1, Score: 99}
	fmt.Println(student)

	person2 := entity.Person{Name: "冰冰老师", Age: 30}
	teacher := entity.Teacher{Person: person2}
	teacher.Email = "email"
	// 继承的字段，可以用 Person.Name 来访问
	teacher.Person.Name = "小红老师"
	// 继承的字段，也可以直接 访问 Age
	teacher.Age = 28
	fmt.Println(teacher)

	// 可以调用继承的方法
	student.Say()
	// 可以调用继承重写的方法
	teacher.Person.Say() // 显式调用父类的方法
	teacher.Say()        // 根据就近原则，调用重写的方法
}
