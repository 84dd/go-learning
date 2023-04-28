package main

import (
	"fmt"
	"io"
	"os"
)

const demoFile = "demo_file.txt"

func main() {
	// 判断文件是否存在
	stat, err := os.Stat(demoFile)
	if err == nil {
		fmt.Println("文件存在，文件名为：", stat.Name())
	} else if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else {
		fmt.Println("未知错误")
	}

	openFile()
	readFile()
	createFile()
}

// 简单打开文件
func openFile() {
	// 用open打开的文件，只能读不能写
	file, err := os.Open(demoFile)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(file.Name())
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			fmt.Println(closeErr)
		}
	}()
}

// 读文件
func readFile() {
	// 用open打开的文件，只能读不能写
	file, err := os.Open(demoFile)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(file.Name())
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			fmt.Println(closeErr)
		}
	}()

	// 方式一：读取固定字节的内容（注意：文件结束符 \n 也会被读取进来）
	// 弊端：这里指定了50字节，但是如果存在中文（3字节），可能会进行截断，导致最后一个字符乱码
	buff1 := make([]byte, 50)
	count1, err1 := file.Read(buff1)
	if err1 != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("-------方式一------")
	fmt.Println("读取到的数量：", count1)
	fmt.Println("读取到的内容：\n", string(buff1))

	// 如果方式一 和 方式二想同时测试，要用 Seek 函数将读指针指向0位置
	// 官方注释：Seek设置下一次读/写的位置
	_, _ = file.Seek(0, 0)

	// 方式二：循环读取
	buff2 := make([]byte, 50)
	contents := make([]byte, 10)
	// 死循环读取
	for {
		count2, err2 := file.Read(buff2)
		// buff2[:count2]  因为最后一次读取到的数量，不一定满
		contents = append(contents[:], buff2[:count2]...)
		if err2 == io.EOF {
			break
		}
	}
	fmt.Println("\n-------方式二------")
	fmt.Println("读取到的内容：\n", string(contents))

	// 还有各种方式，这里不赘述，网上找工具即可
	// 主要方式有以下几种：
	// 1、按块（byte）读取：注意乱码问题
	// 2、按行读取：bufio.NewReader：默然一行长度为4096，也要注意乱码问题
	// 3、整个文件读取：ioutil.ReadFile：不适合大文件
	// 4、按单词读取 bufio.NewScanner

}

// 创建文件
func createFile() {
	// 用open打开的文件，只能读不能写
	file, err := os.Create("write_file.txt")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(file.Name())
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			fmt.Println(closeErr)
		}
	}()

	_, _ = file.Write([]byte{'H', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'd', '\n'})
	_, _ = file.WriteString("你好世界")
}
