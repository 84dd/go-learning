package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 定义一把互斥锁
var lock sync.Mutex

// Sing2 唱歌：每秒停顿一次，总耗时5秒
func Sing2(name string) {
	// 获取锁
	lock.Lock()
	fmt.Println(name)
	for i := 0; i < 5; i++ {
		if i == 2 {
			// 释放锁
			lock.Unlock()
			runtime.Gosched()
		}
		fmt.Printf("%v在唱歌%v，时间 %v\n", name, i, time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}
}

func main() {
	go Sing2("小明")
	go Sing2("Lucy")

	// 主线程睡眠10秒，避免主线程退出，协成也退出
	time.Sleep(time.Second * 10)
}
