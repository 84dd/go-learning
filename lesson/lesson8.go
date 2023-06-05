package lesson

// 参考amos分享的《Go 语言圣经》第五章《函数》匿名函数/可变参数/panic/defer/recover 的基础知识，完成以下练习题，在你个人目录下创建 lesson8.go
// 1.编写一个函数 adder，该函数返回一个匿名函数，该匿名函数接受一个 int 类型的参数， 并将其与之前传递给 adder 的所有参数相加，并返回累加的结果。
// 例如，adder(1) 返回一个函数，该函数接受一个 int 类型的参数 x，并返回 1+x；
// 再次调用 adder(1)，返回的函数将接受一个 int 类型的参数 y，并返回 1+x+y；以此类推。

func Adder(x int) func(int) int {
	// 记录历史值
	var history = x
	return func(x int) int {
		// 累加
		history = history + x
		return history
	}
}

// 2.编写一个函数 divide，该函数接受两个 int 类型的参数 a 和 b，并返回 a 除以 b 的结果 int。
// 但是，在 b 等于 0 时，该函数会引发一个 panic。
// 编写另外一个函数 doDivide，无需返回参数，该函数调用 divide 函数并使用 recover 恢复 panic，以确保程序不会崩溃。
// 对于部分同学纠结 doDivide 传参问题，doDivide() 和 doDivide(a, b int) 都可以

func Divide(a, b int) int {
	// 利用defer特性，捕获异常
	// 注意：一定要在 panic 前，定义 defer
	defer DoDivide()

	if b == 0 {
		panic("除数不能为0")
	}
	return a / b
}

// 二选一
func DoDivide() {
	// 利用 recover() 获取异常，当异常不为空，进入判断
	if err := recover(); err != nil {
		// 捕获了异常，程序恢复
	}
}
