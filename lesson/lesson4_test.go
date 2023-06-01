package lesson

import (
	"fmt"
	"testing"
)

func TestHasPrefixAndSuffix(t *testing.T) {
	fmt.Println(HasPrefixAndSuffix("xxxabcyyy", "xxx", "yyy"))
	fmt.Println(HasPrefixAndSuffix("xxxabcyyy", "xxxxxxxxxxxxxxxxx", "yyy"))
	fmt.Println(HasPrefixAndSuffix("xxxabcyyy", "xxx", "yyyyyyyyyyyyyyyyyyyy"))
	fmt.Println(HasPrefixAndSuffix("xxxabcyyy", "xxx11", "11yyy"))
}

func TestLast(t *testing.T) {
	fmt.Println(Last("xxxabcyyy", 'a'))
	fmt.Println(Last("xxxabcyyy", 'q'))
	fmt.Println(Last("xxxabcyyy", 'y'))
	fmt.Println(Last("xxxabcyyy", 'x'))
}
