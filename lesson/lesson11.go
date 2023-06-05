package lesson

import (
	"bytes"
	"fmt"
	"text/template"
)

// 参考 fanny 分享的《Go 语言圣经》第四章《复合数据类型》JSON/文本和HTML模板
// 完成以下练习题，在你个人目录下创建 lesson11.go
// 1.现在需要你使用 go 语言的结构体组合来设计测试中心的组织架构，请至少包含 QualityCenter - 测试中心, Team - 小组
// 以及 Member - 成员这三种结构体，在成员属性中，由于年龄 Age 属于敏感信息，序列化时需要忽略

type Member struct {
	ID   int    // 主键
	Name string // 姓名
	Age  int8   `json:"-"` // 年龄，敏感信息，序列化时需要忽略
	team Team   // 属于某小组
}

type Team struct {
	ID          int           // 主键
	Name        string        // 小组名称
	MemberCount int           // 小组有多少城市人员
	center      QualityCenter // 属于某测试中心
}

type QualityCenter struct {
	ID        int    // 主键
	Name      string // 测试中心名称
	TeamCount int    // 测试用心有多少小组
}

// 2.给定一个 JSON 字符串，要求实现一个函数，将 JSON 字符串解析为结构体对象。
// 接着，使用 text/template 包，根据给定的模板字符串 tmplStr 和解析出的结构体对象，生成一个新的字符串。
// json 字符串举例
// jsonStr := `{
// 		"id": 1,
// 		"name": "Laptop",
// 		"price": 999.99
// }`

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func ProductTemplate(s string) string {
	// 反序列化 JSON 字符串为 Product 结构体
	// 定义模板字符串
	tmplStr := "Product Information:\nID: {{.ID}}\nName: {{.Name}}\nPrice: {{.Price}}"
	// 使用 text/template 包，根据模板字符串和 Product 结构体对象，生成结果字符串，再将结果字符串打印并返回

	// 初始化，解析
	tmpl, err := template.New("test").Parse(s)
	if err != nil {
		panic(err)
	}

	// 申请结果缓冲区
	buf := new(bytes.Buffer)
	// 根据模版，将结果写入缓冲区
	err = tmpl.Execute(buf, tmplStr)
	if err != nil {
		panic(err)
	}

	// 缓冲区转字符串
	res := buf.String()
	// 输出
	fmt.Println(res)
	// 返回
	return res
}
