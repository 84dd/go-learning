package lesson

// 参考Erin分享的《Go 语言圣经》第四章《复合数据类型》map/结构体 的基础知识，完成以下练习题，在你个人目录下创建 lesson6.go
// 1.请编写一个函数 ReverseMap，该函数收一个 map[string]string 类型的值，
// 返回一个新的 map[string]string 。该函数的作用是将输入的 map 中的键值对反转，并返回反转后的结果。
// 假设入参 map m 中存在相同的 value，反转时将 key 用英文逗号分隔符 , 拼接起来，无需考虑 k/v 为空字符的情况
// 例如 m := map[string]string { "apple": "red", "grape": "red" }
// 需要返回 { "red": "apple,grape" }

func ReverseMap(m map[string]string) map[string]string {
	r := make(map[string]string)
	for k, v := range m {
		rv, has := r[v]
		if has {
			k = rv + "," + k
		}
		r[v] = k
	}
	return r
}

// 2.请定义一个名为 Rectangle 的结构体，该结构体包含以下字段：
// Width，表示矩形的宽度，类型为 int。
// Height，表示矩形的高度，类型为 int。
// 请编写一个名为 Area 的方法，该方法接受一个 Rectangle 实例，并返回该矩形的面积，类型为 int （忽略溢出） 。

type Rectangle struct {
	Width  int
	Height int
}

func (r *Rectangle) Area() int {
	return r.Width * r.Height
}
