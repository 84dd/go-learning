package lesson

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
	a1 := [10]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var n1 int8 = 2
	fmt.Println(Find(a1, n1))

	a2 := [10]int32{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	var n2 int8 = 6
	fmt.Println(Find(a2, n2))
}

func TestCommon(t *testing.T) {
	s1 := []string{"aaa", "bbb", "ccc", "ddd"}
	s2 := []string{"111", "bbb", "333", "ccc", "333"}
	fmt.Println(Common2(s1, s2))
}
