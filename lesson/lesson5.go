package lesson

import "fmt"

// 参考Erin分享的《Go 语言圣经》第四章《复合数据类型》数组/切片 的基础知识，完成以下内容
// 在你个人目录下创建 lesson5.go，完成如下函数

// 1.在长度为 10 的 int32 数组中，找到第 n 大的数， 1=< n <= 10（代码里不用考虑 n 的其他情况）；
// 以长度 5 的数组举例：[1, 6, 3, 5, 8]，传入 n=1（第一大） 返回 8，传入 n=5（第五大） 返回 1。
// https://blog.csdn.net/weixin_42711189/article/details/116351162
func Find(a [10]int32, n int8) int32 {
	// 可以实现不同的排序算法
	bubbleSort(&a)
	fmt.Println(a)
	return a[len(a)-int(n)]
}

// 方法1：冒泡排序
func bubbleSort(a *[10]int32) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

// 2.两个包含字符串的切片 s1/s2，找到这两个切片中的公共字符串，如果有重复的也要找出来。
// 然后，将所有的公共字符串存储在一个新的切片 s3 中，并返回。
func Common(s1, s2 []string) []string {
	s3 := make([]string, 0)
	// key=切片中的字符串  value=是否出现过
	m := make(map[string]bool)
	// 由于要找重复值，所以先连接两个切片
	s := append(s1, s2...)
	for _, v := range s {
		if m[v] {
			// 如果出现过，则放在结果集中
			s3 = append(s3, v)
		}
		// 每一次遍历，都设置为true
		m[v] = true
	}
	return s3
}

// 3.思考题（可选题）：第 2 题中，如果切片s1/s2的 size 非常大，函数如何改造可以提高其性能。
// 回答：
// 首先要知道一点：切片的长度可以随着元素数量的增加而增长, 但不会随着元素减少而减小
// 所以 s := append(s1, s2...) 这行会占用极大的空间，随着for不断循环，排除的数据越来越多，但是切片依然不会变小
// 优化1：不append，分两次遍历，具体看 Common1 方法

func Common1(s1, s2 []string) []string {
	s3 := make([]string, 0)
	// key=切片中的字符串  value=是否出现过
	m := make(map[string]bool)
	for _, v := range s1 {
		if m[v] {
			// 如果出现过，则放在结果集中
			s3 = append(s3, v)
		}
		// 每一次遍历，都设置为true
		m[v] = true
	}
	for _, v := range s2 {
		if m[v] {
			// 如果出现过，则放在结果集中
			s3 = append(s3, v)
		}
		// 每一次遍历，都设置为true
		m[v] = true
	}
	return s3
}

// 优化2： append 函数会在需要时自动扩展切片的容量，频繁的扩容操作会带来较大的性能开销，因此我们可以在使用 append 函数前预分配切片的容量，以减少扩容操作的次数
// 我们可以为s3切片预分配容量为：最小切片的一半长度，具体看 Common2 方法
func Common2(s1, s2 []string) []string {
	len1 := len(s1)
	len2 := len(s2)
	var len3 int
	if len1 > len2 {
		len3 = len2
	} else {
		len3 = len1
	}
	s3 := make([]string, len3/2)

	// key=切片中的字符串  value=是否出现过
	m := make(map[string]bool)
	for _, v := range s1 {
		if m[v] {
			// 如果出现过，则放在结果集中
			s3 = append(s3, v)
		}
		// 每一次遍历，都设置为true
		m[v] = true
	}
	for _, v := range s2 {
		if m[v] {
			// 如果出现过，则放在结果集中
			s3 = append(s3, v)
		}
		// 每一次遍历，都设置为true
		m[v] = true
	}
	// 取切片的有效部分
	return s3[:0]
}
