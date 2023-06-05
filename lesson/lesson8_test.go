package lesson

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	addOne := Adder(1)
	fmt.Println(addOne(2)) // 输出 3 = 1 + 2
	fmt.Println(addOne(3)) // 输出 6 = 1 + 2 + 3

	addTwo := Adder(2)
	fmt.Println(addTwo(4)) // 输出 6 = 2 + 4
	fmt.Println(addTwo(8)) // 输出 14 = 2 + 4 + 8
}

func TestDivide(t *testing.T) {
	fmt.Println(Divide(10, 0))
	fmt.Println(Divide(10, 2))
	fmt.Println(Divide(10, 0))
	fmt.Println(Divide(10, 5))
}
