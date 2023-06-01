package main

import "fmt"

// Shape https://blog.csdn.net/pengxianchen/article/details/125711079
type Shape interface {
	// GetArea 在接口中定义方法，如果是空参，表示可以接收任意参数
	// 也可以指定参数类型和返回类型
	GetArea() float64
	GetPerimeter() float64
}

// Circle 定义结构体
type Circle struct {
	R float64
}

// GetArea Circle实现方法
func (c Circle) GetArea() float64 {
	fmt.Println("enter Circle.GetArea")
	return 3.14 * c.R * c.R
}

// GetPerimeter Circle实现方法
func (c Circle) GetPerimeter() float64 {
	return 3.14 * 2 * c.R
}

type Square struct {
	W float64
	H float64
}

func (s Square) GetArea() float64 {
	fmt.Println("enter Square.GetArea")
	return s.W * s.H
}

func (s Square) GetPerimeter() float64 {
	return 2 * (s.W + s.H)
}

func main() {
	var s Shape

	s = Circle{10}
	s.GetArea() // 输出：enter Circle.GetArea

	s = Square{10, 10}
	s.GetArea() // 输出：enter Square.GetArea

	// 利用ok-idiom模式将接口类型还原为原始类型
	// s.(Square)这种格式我们称之为: 类型断言
	r1, ok1 := s.(Circle)
	fmt.Println(r1, ok1)

	if r2, ok2 := s.(Square); ok2 {
		fmt.Println("参数 W = ", r2.W)
		fmt.Println("参数 H = ", r2.H)
	}
}
