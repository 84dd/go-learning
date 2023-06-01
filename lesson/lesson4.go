package lesson

// 参考浩天分享的《Go 语言圣经》第三章《基础数据类型》的基础知识，完成以下内容
// 在你个人目录下创建 lesson4.go，完成如下函数

// 1.判断字符串 s 是否包含某个字符串前缀 prefix 和某个字符串后缀 suffix，都包含返回 true 否则 false， 允许 prefix 或 suffix 其一为空，若同时为空返回 false
func HasPrefixAndSuffix(s, prefix, suffix string) bool {
	if len(s) == 0 {
		return false
	}
	// 判断前缀
	hasPrefix := false
	if len(prefix) > 0 && len(prefix) <= len(s) {
		hasPrefix = s[0:len(prefix)] == prefix
	}
	if len(prefix) > 0 && !hasPrefix {
		// 如果前缀有长度并且不包含，可以提前知道结果并返回
		return false
	}

	// 判断后缀
	hasSuffix := false
	if len(suffix) > 0 && len(prefix) <= len(s) {
		hasSuffix = s[len(s)-len(prefix):] == suffix
	}
	return hasPrefix && hasSuffix
}

// 2.给出一串字符串 s（长度一万内），计算某个字符 b 在该字符串中最后出现的索引位置，返回为有符号整数，没有找到返回 -1；
func Last(s string, b byte) int {
	if len(s) == 0 {
		return -1
	}

	// 一般方法：从前往后
	//idx := -1
	//for i := 0; i < len(s); i++ {
	//	if s[i] == b {
	//		idx = i
	//	}
	//}

	// 改进方法：从后往前，抓住重点“最后”
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == b {
			return i
		}
	}

	return -1
}
