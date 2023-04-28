package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Println(now)
	fmt.Println("获取1970-1-1距今的时间（秒）：", now.Unix())
	fmt.Println("获取1970-1-1距今的时间（纳秒）：", now.UnixNano())

	fmt.Println("睡眠前的时间为", time.Now().Format("2006-01-02 15:04:05"))
	// 睡眠一秒
	time.Sleep(time.Second)
	fmt.Println("睡眠后的时间为", time.Now().Format("2006-01-02 15:04:05"))

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	// 格式化方式1
	str1 := fmt.Sprintf("%d-%d-%d %d:%d:%d", year, month, day, hour, minute, second)
	fmt.Println(str1) // 2023-4-28 9:21:50

	// 格式化方式2
	// 2006/01/02 15:04:05这个字符串是Go语言规定的，各个数字都是固定的
	str2 := now.Format("2006/01/02 15:04:05")
	fmt.Println(str2) // 2023/04/28 09:23:21
	str3 := now.Format("2006-01-02 15:04:05")
	fmt.Println(str3) // 2023-04-28 09:23:50
	str4 := now.Format("06-1-2 15:04:5")
	fmt.Println(str4) // 23-4-28 10:17:4

	// 查看手册，为什么go的时间格式化如此诡异
	whyGoDateFormat()
}

// https://blog.csdn.net/asd1126163471/article/details/119583038
func whyGoDateFormat() {
	// 仅记一点：数字就是格式
	// 2006/01/02 15:04:05   相当于java的 yyyy/MM/dd HH:mm:ss  24小时制
	// 2006/01/02 3:04:05   相当于java的 yyyy/MM/dd hh:mm:ss   12小时制
	// 06/1/02 3:4:05   相当于java的 yy/M/dd hh:m:ss

	// 但为什么是这些数字呢？
	// Go 是这么设计的
	// 1: month (January, Jan, 01, etc)
	// 2: day
	// 3: hour (15 is 3pm on a 24 hour clock)
	// 4: minute
	// 5: second
	// 6: year (2006)
	// 7: timezone (GMT-7 is MST)

	//  按 ANSIC 标准的日期格式，
	// 月、日、时、分、秒、年，最后加 MST 时区
	// 对应就是 1 2 3 4 5 6 7。同时还可以随意加星期几。
	// 发现没有？围绕着 1 2 3 4 5 6 7 随意变化，真的不要太爽
}
