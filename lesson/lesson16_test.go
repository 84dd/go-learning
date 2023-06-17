package lesson

import (
	"fmt"
	"testing"
	"time"
)

func TestTaskScheduler(t *testing.T) {
	ts := NewTaskScheduler()

	// 注册任务 1
	ts.RegisterTask("task001", time.Now().Add(3*time.Second), func() { printId("task001") })

	// 注册任务 2
	ts.RegisterTask("task002", time.Now().Add(10*time.Second), func() { printId("task002") })

	// 注册任务 3，但又马上取消
	ts.RegisterTask("task003", time.Now().Add(10*time.Second), func() { printId("task003") })
	ts.CancelTask("task003")

	// 注册任务 4
	ts.RegisterScheduleTask("task004", 3*time.Second, func() { printId("task004") })

	// 注册任务 5
	ts.RegisterScheduleTask("task005", time.Second, func() { printId("task005") })

	// 注册任务 6，但又马上取消
	ts.RegisterScheduleTask("task006", 5*time.Second, func() { printId("task006") })
	ts.CancelTask("task006")

	go ts.Start()

	// 开始任务8秒后，将任务2、4取消，所以任务2永远不会被执行，任务4会被执行两次
	time.Sleep(8 * time.Second)
	ts.CancelTask("task002")
	ts.CancelTask("task004")

	// 再过10秒后，停止任务调度器
	time.Sleep(10 * time.Second)
	ts.Stop()
	ts.wg.Wait() // 等待
}

// 打印id
func printId(id string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "执行任务-->>>", id)
}
