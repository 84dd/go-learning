package main

import (
	"fmt"
	"time"
)

func main() {
	demoTimer()
	demoTicker()
}

// 一次性延时，可以理解为js中的 setTimeout
func demoTimer() {
	fmt.Println("开始时间：", time.Now().Format("15:04:05.000"))
	timer := time.NewTimer(time.Second * 3)
	fmt.Println("读取前不会被阻塞", time.Now().Format("15:04:05.000"))
	time.Sleep(time.Second * 5)
	fmt.Println("再延时5秒后，依然不阻塞", time.Now().Format("15:04:05.000"))
	fmt.Println("读取前的时间：", time.Now().Format("15:04:05.000"))
	end := <-timer.C
	fmt.Println("结束时间1：", time.Now().Format("15:04:05.000"))
	fmt.Println("结束时间2：", end.Format("15:04:05.000"))

	/*
		// 上面的输出如下
		开始时间： 15:04:41.526
		读取前不会被阻塞 15:04:41.534
		再延时5秒后，依然不阻塞 15:04:46.542    // 这个延时，是针对 开始时间 的，所以比开始时间多5秒
		读取前的时间： 15:04:46.542
		结束时间1： 15:04:46.542    			// 调用end := <-timer.C分两种情况
											// 1：如果当前时间已经满足NewTimer的延时，则不阻塞
											// 2：如果当前时间还没到NewTimer的延时，则再阻塞剩余的时间，
											//    将time.Sleep(time.Second * 5)改为1秒看效果
		结束时间2： 15:04:44.539				// 这个时间是针对time.NewTimer(time.Second * 3)后的延时时间
	*/
}

// 周期性延时，可以理解为js中的 setInterval
func demoTicker() {
	ticker := time.NewTicker(time.Second)
	for i := 0; i < 10; i++ {
		_ = <-ticker.C
		fmt.Printf("第%v次，时间：%v\n", i, time.Now().Format("15:04:05.000"))
	}
}
