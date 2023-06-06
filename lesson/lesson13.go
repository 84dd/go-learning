package lesson

import (
	"bytes"
	"strings"
)

// 参考 erin 分享的《Go 语言圣经》第十一章《测试》测试函数/基准测试/数据驱动测试
// 在你个人目录下创建 lesson13.go 以及 lesson13_test.go ，然后完成以下练习题：
// 1.编写函数 IsPalindrome 判断字符串是否是回文(正反读完全一样)，并为其编写测试函数，请至少包含 10 条测试数据
// •英文大写小写视为相等
// •空字符视为回文
// •字符串只有一个字符视为回文
// •需要考虑字符串中包含特殊符号的情况
// 是回文：
// 		1.上海自来水来自海上
// 		2.Madam
// 不是回文：
// 		1.你好吗
// 		2.hello

func IsPalindrome(s string) bool {
	// 考点1：字符串转小写
	// 考点2：考虑中文（rune）
	s1 := []rune(strings.ToLower(s))
	size := len(s1)
	if size < 2 {
		return true
	}

	// 考点3：回文的特性
	middle := size / 2
	for i := 0; i < middle; i++ {
		if s1[i] != s1[size-1-i] {
			return false
		}
	}
	return true
}

// 2.编写两个（或更多）函数 ConcatStringsX，该函数中字符切片根据某个特定分隔符做拼接返回新的字符串，
// 比如 ["a", "b", "c"]，分隔符为 _, 则返回为 a_b_c。
// 并给这两个（或更多）函数编写 基准测试 用例，切片的长度可以考虑 10，100，1000，10000 的情况
func ConcatStrings1(s []string, sep string) string {
	// 调用封装好的函数
	return strings.Join(s, sep)
}

func ConcatStrings2(s []string, sep string) string {
	if len(s) < 2 {
		return s[0]
	}
	// 自己写for（性能肯定不好）
	str := ""
	// 小技巧：步进到 len(s)-2 即可
	for i := 0; i < len(s)-2; i++ {
		str += s[i] + sep
	}
	str += s[len(s)-1]
	return str
}

func ConcatStrings3(s []string, sep string) string {
	if len(s) < 2 {
		return s[0]
	}
	// 创建一个空的缓冲区
	var buf bytes.Buffer
	// 小技巧：步进到 len(s)-2 即可
	for i := 0; i < len(s)-2; i++ {
		buf.WriteString(s[i])
		buf.WriteString(sep)
	}
	buf.WriteString(s[len(s)-1])
	return buf.String()
}
