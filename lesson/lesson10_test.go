package lesson

import (
	"fmt"
	"testing"
)

func TestBoook(t *testing.T) {
	book1 := Book{1, "Go从入门到放弃", "高老师", "7-309-04547", 100}
	book1.Borrow("小明")

	book2 := Book{2, "如何有效治疗颈椎病", "何医生", "7-309-04999", 990}
	book2.Borrow("小明")

	book3 := Book{3, "如何在幼儿园开展编程教学", "胡教授", "6-309-0358", 120}
	book3.Borrow("小红")

	fmt.Println(ListBooks("小明"))
	fmt.Println(ListBooks("小红"))
}
