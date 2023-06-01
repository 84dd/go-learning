package lesson

import (
	"fmt"
	"testing"
)

func TestLcm(t *testing.T) {
	num1 := Lcm(4, 6)
	fmt.Println(num1) // 12

	num2 := Lcm(3, 5)
	fmt.Println(num2) // 15

	// 两数相等的情况
	num3 := Lcm(7, 7)
	fmt.Println(num3) // 7

	// 一个数刚好是另一个数的倍数
	num4 := Lcm(3, 12)
	fmt.Println(num4) // 12

	// 一个数刚好是另一个数的倍数
	num5 := Lcm(12, 4)
	fmt.Println(num5) // 12
}
