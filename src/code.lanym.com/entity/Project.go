package entity

// Project 演示匿名属性
type Project struct {
	Student         // 可以匿名，属性名默认是类型名
	Stu2    Student // 显式属性名
	Status  int
	Remark  string
}
