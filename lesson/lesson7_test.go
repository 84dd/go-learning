package lesson

import (
	"fmt"
	"testing"
)

func TestFactorial(t *testing.T) {
	fmt.Println(Factorial(1)) // 1
	fmt.Println(Factorial(2)) // 2
	fmt.Println(Factorial(3)) // 6
	fmt.Println(Factorial(4)) // 24
}

func TestSplit(t *testing.T) {
	n := []int{1, 2, 3, 4, 5}
	f := func(x int) bool { return x%2 == 0 }
	fmt.Println(Split(n, f))
}
