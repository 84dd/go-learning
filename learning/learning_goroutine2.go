package main

import (
	"fmt"
	"go-learning/src/code.lanym.com/mygogo"
	"runtime"
	"time"
)

func main() {
	// 该方法有 礼让协成
	go mygogo.Sing1()
	// 该方法有 退出协成
	go mygogo.Dance1()

	// 主线程睡眠10秒，避免主线程退出，协成也退出
	time.Sleep(time.Second * 6)

	cpuNum := runtime.NumCPU()
	fmt.Println("cpu数量：", cpuNum)
}
