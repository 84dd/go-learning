package lesson

import (
	"fmt"
	"testing"
	"time"
)

func TestTaskScheduler(t *testing.T) {
	ts := NewTaskScheduler()
	fmt.Println(ts)
	defer ts.Stop()
	ts.Start()

	// 注册任务 1，设置执行时间为当前时间 + 5 秒，执行函数为打印任务 ID
	ts.RegisterTask(
		"task001",
		time.Now().Add(5*time.Second),
		func() { fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " task001") },
	)

	// 注册任务 2，设置执行时间为当前时间 + 10 秒，执行函数为打印任务 ID
	ts.RegisterTask(
		"task002",
		time.Now().Add(10*time.Second),
		func() { fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " task002") },
	)

	// 注册任务 3，但又马上取消
	ts.RegisterTask(
		"task003",
		time.Now().Add(10*time.Second),
		func() { fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " task003") },
	)
	ts.CancelTask("task003")

	// 等待 15 秒，确保任务调度器有足够的时间执行任务
	time.Sleep(15 * time.Second)

	// 注册任务 4，设置执行间隔为 5 秒，执行函数为打印任务 ID
	ts.RegisterScheduleTask(
		"task004",
		5*time.Second,
		func() { fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " task004") },
	)

	// 注册任务 5，但又马上取消
	ts.RegisterScheduleTask(
		"task005",
		5*time.Second,
		func() { fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " task005") },
	)
	ts.CancelTask("task005")

	// 等待一段时间，确保任务调度器有足够的时间执行任务
	time.Sleep(60 * time.Second)
}
