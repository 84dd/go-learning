package main

import (
	"go-learning/src/code.lanym.com/mygogo"
	"time"
)

func main() {
	// 开启一个协成
	go mygogo.Sing()
	go mygogo.Dance()

	// 主线程睡眠10秒，避免主线程退出，协成也退出
	time.Sleep(time.Second * 10)
}
