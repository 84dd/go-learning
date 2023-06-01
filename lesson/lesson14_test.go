package lesson

import (
	"fmt"
	"testing"
)

func TestFileNameSearchest(t *testing.T) {
	var dirPath string // 搜索的目录路径
	var keyword string // 搜索的关键字
	// 情况1：路径不存在，返回 []
	dirPath = "D:\\work\\IdeaProjects\\xxxxxxxxx\\src\\main"
	keyword = ".java"
	fmt.Println(FileNameSearch(dirPath, keyword))

	// 情况2：关键字为空
	dirPath = "D:\\work\\IdeaProjects\\drawio\\src\\main"
	keyword = ""
	fmt.Println(FileNameSearch(dirPath, keyword))

	// 情况3：关键字不存在
	dirPath = "D:\\work\\IdeaProjects\\drawio\\src\\main"
	keyword = ".xxxxxxxxxxx"
	fmt.Println(FileNameSearch(dirPath, keyword))

	// 情况4：正常
	dirPath = "D:\\work\\IdeaProjects\\drawio\\src\\main"
	keyword = ".java"
	fmt.Println(FileNameSearch(dirPath, keyword))
}
