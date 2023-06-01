package lesson

import (
	"math"
)

// 参考财源分享的《Go 语言圣经》第二章《程序结构》基础知识，完成以下内容
// 在你个人目录下创建 lesson3.go，完成如下方法/函数后
// 1.计算两个整数值 int32（x, y）的最小公倍数 z（LCM），函数名为：lcm，假设 y != 0，且结果不会超过 int32 最大值
// 2.通过第一课里面的 main.go 文件的 main 方法里面调用对应方法/函数
// 概念：两个或多个整数公有的倍数叫做它们的公倍数，最小公倍数是指所有公倍数中除0以外最小的一个公倍数。
// 例如：4 和 6 的最小公倍数就是12   （4*3=12，6*2=12）
// 例如：5 和 3 的最小公倍数就是15   （5*3=15，3*5=15）

func Lcm(x, y int32) int32 {
	if x == y {
		// 两数相等，则该数就是最小公倍数
		return x
	}

	// 先比较两数，并确定x为小数，y为大数
	if x > y {
		x, y = y, x
	}

	// 定义倍数
	var idx int32 = 0
	for {
		// 每次循环，倍数 + 1
		idx += 1
		// 获取最小公倍数的思路：
		// 1、取 大数除小数的余数，如果余数为0，则 大数乘以当前倍数，就是最小公倍数
		// 2、如果余数不为0，倍数+1，继续判断
		if math.Remainder(float64(y*idx), float64(x)) == 0 {
			break
		}
	}
	return y * idx
}