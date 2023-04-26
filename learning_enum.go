package main

import (
	"fmt"
	"go-learning/src/code.lanym.com/comment"
)

func main() {
	// 枚举，利用iota的特性，按行加1
	const (
		type1 = iota
		type2
		type3
		type4
	)

	fmt.Println("jpg 数字标识：", comment.Jpg)
	fmt.Println("png 数字标识：", comment.Png)
	fmt.Println("gif 数字标识：", comment.Gif)

	fmt.Println("mp4 后缀：", comment.Mp4)
	fmt.Println("mp3 后缀：", comment.Mp3)
}
