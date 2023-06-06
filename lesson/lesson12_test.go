package lesson

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	fmt.Println(Calculate(Addition{}, 1, 2))
	fmt.Println(Calculate(Subtraction{}, 3, 4))
	fmt.Println(Calculate(Multiplication{}, 5, 6))
	fmt.Println(Calculate(Division{}, 7, 8))
	fmt.Println(Calculate(Division{}, 9, 0))
}

func TestUniversalMap(t *testing.T) {
	slice1 := []int{1, 2, 3}
	mapper1 := func(value interface{}) interface{} {
		return value.(int) * value.(int)
	}
	res1 := UniversalMap(slice1, mapper1)
	fmt.Println("平方(int)：", res1)

	slice2 := []float64{1.1, 2.2, 3.3}
	mapper2 := func(value interface{}) interface{} {
		return value.(float64) * value.(float64)
	}
	res2 := UniversalMap(slice2, mapper2)
	fmt.Println("立方(float64)：", res2)

	slice3 := []string{"abcabc", "qwert123456", ""}
	mapper3 := func(value interface{}) interface{} {
		return len(value.(string))
	}
	res3 := UniversalMap(slice3, mapper3)
	fmt.Println("取长度(string)：", res3)

	slice4 := []float32{1.1, 2.2, 3.3}
	mapper4 := func(value interface{}) interface{} {
		return value.(float32) * value.(float32)
	}
	res4 := UniversalMap(slice4, mapper4)
	fmt.Println("取长度(接口明确不支持float32，会返回nil)：", res4)
}
