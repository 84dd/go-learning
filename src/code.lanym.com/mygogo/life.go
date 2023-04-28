package mygogo

import (
	"fmt"
	"runtime"
	"time"
)

// Sing 唱歌：每秒停顿一次，总耗时5秒
func Sing() {
	for i := 0; i < 5; i++ {
		fmt.Printf("我在唱歌%v，时间 %v\n", i, time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}
}

// Dance 跳舞：每秒停顿一次，总耗时5秒
func Dance() {
	for i := 0; i < 5; i++ {
		fmt.Printf("dancing %v，time %v\n", i, time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}
}

// Sing1 唱歌：每秒停顿一次，总耗时5秒
func Sing1() {
	for i := 0; i < 5; i++ {
		if i == 1 {
			// Gosched 表示使当前go程放弃处理器，让其他go程运行
			// 相当于java中的 yield 礼让
			runtime.Gosched()
		}
		fmt.Printf("我在唱歌%v，时间 %v\n", i, time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}
}

// Dance1 跳舞：每秒停顿一次，总耗时5秒
func Dance1() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			// Goexit 退出当前协成
			runtime.Goexit()
		}
		fmt.Printf("dancing %v，time %v\n", i, time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}
}
