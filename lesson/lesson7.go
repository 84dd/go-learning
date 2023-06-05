package lesson

// 参考乾航分享的《Go 语言圣经》第五章《函数》函数声明/递归函数/函数值 的基础知识，完成以下练习题，在你个人目录下创建 lesson7.go
// 1.编写一个递归函数，计算一个正整数的阶乘。不用考虑入参为负数的情况。

func Factorial(n int) int {
	// 编写你的代码
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// 2.编写一个函数，该函数接受一个整数切片n和一个函数f作为参数，然后根据函数f的返回值将整数切片拆分成两个切片。
// 函数f的入参是切片n的元素，返回值是一个布尔值，指示该元素是否属于第一个切片。
// 如果返回 true，则将该元素添加到第一个切片中；否则，将该元素添加到第二个切片中。
// 例如，如果输入整数切片是 [1, 2, 3, 4, 5]，函数是 func(x int) bool { return x % 2 == 0 }，
// 那么函数应该返回两个整数切片：[2, 4] 和 [1, 3, 5]。

// 传给该函数的函数形参f由你自行定义，满足上面的要求即可
func Split(n []int, f func(int) bool) ([]int, []int) {
	// 考点1：range 的使用，或者切片的遍历
	// 考点2：append 的使用，或者切片添加数据
	// 考点3：函数入参
	// 考点4：多返回
	var n1, n2 []int
	for _, num := range n {
		if f(num) {
			n1 = append(n1, num)
		} else {
			n2 = append(n2, num)
		}
	}
	return n1, n2
}
