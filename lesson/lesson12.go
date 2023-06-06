package lesson

import "fmt"

// 参考 浩天 分享的《Go 语言圣经》第七章《接口》接口定义/接口类型/接口值/空接口
// 在你个人目录下创建 lesson12.go ，然后完成以下练习题：
// 1.结合第二课的题目（加减乘除）以及本周的课程内容，通过接口的方式来实现一个简单的计算器程序。
// 定义一个名为Operation的接口，包含 Execute() 的方法，方法接受两个浮点数作为参数，返回一个浮点数。
// 创建四个结构体Addition、Subtraction、Multiplication和Division， 分别表示加/减/乘/除操作，并实现Operation接口。
// 然后，编写一个函数 Calculate(o Operation, x, y float64) float64，接收一个操作及两个操作数，并返回计算结果。
// •无需考虑数据溢出的情况
// •暂时不用考虑除数为 0 的异常情况

type Operation interface {
	Execute(x, y float64) float64
}

type Addition struct {
}

func (_ Addition) Execute(x, y float64) float64 {
	return x + y
}

type Subtraction struct {
}

func (_ Subtraction) Execute(x, y float64) float64 {
	return x - y
}

type Multiplication struct {
}

func (_ Multiplication) Execute(x, y float64) float64 {
	return x * y
}

type Division struct {
}

func (_ Division) Execute(x, y float64) float64 {
	defer func() {
		if err := recover(); err != nil {
			// 捕获了异常，程序恢复
			fmt.Println("panic异常：", err)
		}
	}()
	if y == 0 {
		panic("除数不能为0")
	}
	return x / y
}

func Calculate(o Operation, x, y float64) float64 {
	return o.Execute(x, y)
}

// 2.定义一个映射函数 UniversalMap，它可以将切片里面的元素通过映射函数 mapper 全部转换成新的元素。
// 函数 UniversalMap 接受一个切片和一个函数值 mapper（该函数的入参为切片元素，返回一个新的元素），返回一个新的切片。
// 使用类型断言或类型开关（type switch）来确定传入切片的实际类型，并分别处理不同类型的切片数据。
// 请为以下类型的切片编写映射处理代码：[]int、[]float64、[]string，其他类型一律返回 nil
// 举例：
// 当传入的 slice = int[1, 2, 3] 时，定义的 mapper := func(value interface{}) interface{} {
// return value.(int) * value.(int)
// }
// 则 UniversalMap 返回 int[1, 4, 9]

// interface {} 或者 any 均可
func UniversalMap(slice interface{}, mapper func(interface{}) interface{}) interface{} {
	// 实现类型判断和映射逻辑
	switch slice.(type) {
	case []int:
		// 考点1：类型转换
		s1 := slice.([]int)
		// 考点2：要提前make分配空间
		res := make([]interface{}, len(s1))
		for i := 0; i < len(s1); i++ {
			res[i] = mapper(s1[i])
		}
		return res
	case []float64:
		s1 := slice.([]float64)
		res := make([]interface{}, len(s1))
		for i := 0; i < len(s1); i++ {
			res[i] = mapper(s1[i])
		}
		return res
	case []string:
		s1 := slice.([]string)
		res := make([]interface{}, len(s1))
		for i := 0; i < len(s1); i++ {
			res[i] = mapper(s1[i])
		}
		return res
	}
	return nil
}
