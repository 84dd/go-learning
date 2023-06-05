package lesson

// 参考 katala 分享的《Go 语言圣经》第六章《方法》嵌入结构体来拓展类型/方法值和方法表达式/封装
// 完成以下练习题，在你个人目录下创建 lesson10.go
// 设计一个程序来管理图书馆的书籍。该程序应包含一个名为 Book 的结构体类型，其中包含以下字段：
// ID：表示书籍的唯一标识符。 Title：表示书籍的标题。 Author：表示书籍的作者。 ISBN：表示书籍的国际标准书号。 Pages：表示书籍的页数。
// 为了拓展 Book 类型，你需要创建一个名为 AudioBook 的嵌入式结构体类型，并为其添加以下字段：
// Narrator：表示有声书的叙述者。 Duration：表示有声书的时长。
// 为了封装 Book 结构体，你应该创建一个名为 Borrow 的方法，用于将书籍借出给某个人。
// 该方法接受一个字符串类型的参数，表示借阅者的姓名，并需要把借书记录保存在 map 中，该 map 用于模拟借书记录。
// 还需要通过另外一个方法 ListBooks 来查询某个借阅者借过的所有书籍编号，没有记录返回空切片即可。

// 书籍借阅记录：
// 初始化map，因为map类型的变量，要make后才能使用，否则报错
var _BorrowRecord = make(map[int]string) // key=book id(int or string)， value=借阅者名字
//var _BorrowRecord map[int]any // 二选一 `any` 可以由你自行定义

type Book struct {
	ID     int    // 表示书籍的唯一标识符。
	Title  string // 表示书籍的标题。
	Author string // 表示书籍的作者。
	ISBN   string // 表示书籍的国际标准书号。
	Pages  int    // 表示书籍的页数。
}

type AudioBook struct {
	Narrator string // 表示有声书的叙述者。
	Duration int    // 表示有声书的时长。
}

// Borrow 把书籍借阅给某个人
func (b *Book) Borrow(s string) {
	// 如果选择 var _BorrowRecord map[int]any 这种方式，可以存储更多的信息
	//row := map[string]any{}
	//row["name"] = s
	//row["title"] = b.Title
	//row["time"] = time.Now().Second()
	_BorrowRecord[b.ID] = s
}

// ListBooks 获取借阅者借的所有书籍编号
func ListBooks(s string) []int {
	var ids []int
	// 遍历
	for id, name := range _BorrowRecord {
		if name == s {
			ids = append(ids, id)
		}
	}
	return ids
}
