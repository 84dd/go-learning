package lesson

import (
	"fmt"
	"testing"
)

func TestReverseMap(t *testing.T) {
	m := map[string]string{"apple": "red", "grape": "red", "banana": "yellow"}
	fmt.Println(ReverseMap(m))
}

func TestArea(t *testing.T) {
	r := Rectangle{2, 2}
	fmt.Println(r.Area())
}
