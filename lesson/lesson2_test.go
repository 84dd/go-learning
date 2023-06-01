package lesson

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	num1 := add(1, 2)
	fmt.Println(num1) // 3

	num2 := sub(3, 4)
	fmt.Println(num2) // -1

	num3 := mult(5, 6)
	fmt.Println(num3) // 30

	num4 := div(7, 8)
	fmt.Println(num4) // 0.88
}
