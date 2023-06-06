package lesson

import (
	"fmt"
	"math/rand"
	"testing"
)

// 单测
func TestIsPalindrome(t *testing.T) {
	fmt.Println(IsPalindrome("上海自来水来自海上"))  // true
	fmt.Println(IsPalindrome("Madam"))      // true
	fmt.Println(IsPalindrome("你好吗"))        // false
	fmt.Println(IsPalindrome("hello"))      // false
	fmt.Println(IsPalindrome("!@#$%%$#@!")) // true
	fmt.Println(IsPalindrome("!@WEDS#@!"))  // false
	fmt.Println(IsPalindrome("a"))          // true
	fmt.Println(IsPalindrome("&"))          // true
	fmt.Println(IsPalindrome("你"))          // true
	fmt.Println(IsPalindrome(""))           // true
}

// 基准测试
func BenchmarkConcatStrings1(b *testing.B) {
	n := rand.Intn(100000)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = "a"
	}
	fmt.Println("切片长度：", len(s))
	fmt.Println("拼接后字符串长度：", len(ConcatStrings1(s, "_")))
}

func BenchmarkConcatStrings2(b *testing.B) {
	s := []string{"a", "b", "c"}
	fmt.Println(s)
	fmt.Println(ConcatStrings1(s, "*"))
}

func BenchmarkConcatStrings3(b *testing.B) {
	n := rand.Intn(100000)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = "a"
	}
	fmt.Println("切片长度：", len(s))
	fmt.Println("拼接后字符串长度：", len(ConcatStrings1(s, "_")))
}
