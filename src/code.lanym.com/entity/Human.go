package entity

import "fmt"

// 注意点1：如果继承的结构体出现重名，则采用就近原则
// 注意点2：结构体可以多重继承
// 注意点3：子类不仅能继承父类的属性，也能继承父类的方法

type Person struct {
	Name string
	Age  int
}

func (p Person) Say() {
	fmt.Println("说话。。。。")
}

type Student struct {
	Person // 相当于java中的继承，但这里是结合，因为go因为结合优于继承
	Score  float32
}

type Teacher struct {
	Person
	Phone, Email string // 也可以定义在同一行
	money        float32
}

func (p Teacher) Say() {
	fmt.Println("老师说话。。。。")
}
