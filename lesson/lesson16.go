package lesson

import (
	"sync"
	"time"
)

// week 23(6/5 - 6/9) 【lesson16】
// 参考 方琦 分享的《Go 语言圣经》第九章《基于共享变量的并发》的内容
// 在你个人目录下创建 lesson16.go ，然后完成以下练习题：
// 实现一个并发安全的任务调度器
// 要求
// 任务调度器，用户调度并发执行的任务；
// 调度器支持如下能力： 注册任务：将任务添加到调度器中，并指定任务的执行时间；将定时任务添加到调度器中，并指定执行间隔
// 并发执行：调度器应该支持并发执行多个任务
// 定时执行：添加的定时任务，可以在设定的时间定时执行
// 取消任务：允许取消已经注册的任务
// 每个任务都有一个唯一标识，并且具有一个需要执行的函数
// 调度器应确保任务的执行顺序和执行时间的准确性
// 调度器应支持并发注册和执行任务的能力，保证并发安全性
// 包含必要的注释和测试案例，以展示任务调度器的功能和并发安全性
// 提示
// stop 通道用作关闭调度器的信号
// sync.Mutex 用来保证并发安全
// sync.WaitGroup 用来保证调度器任务执行完毕，或者已手动关闭
// 调度器运行时建议使用 for select，并且设定一定的描述间隔来检查任务是否需要执行

type TaskScheduler struct {
	tasks map[string]Task
	mu    sync.Mutex
	stop  chan struct{}
	wg    sync.WaitGroup
}

type Task struct {
	// 设置执行时间
	execTime time.Time
	// 设置执行的频率
	interval time.Duration
	f        func()
}

func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks: make(map[string]Task),
		mu:    sync.Mutex{},
		stop:  make(chan struct{}),
		wg:    sync.WaitGroup{},
	}
}

// 注册某个点执行的任务
func (s *TaskScheduler) RegisterTask(id string, t time.Time, f func()) {
	// 你的代码
}

// 注册以某个固定频率执行的任务
func (s *TaskScheduler) RegisterScheduleTask(id string, interval time.Duration, f func()) {
	// 你的代码
}

// 取消任务
func (s *TaskScheduler) CancelTask(id string) {
	// 你的代码
}

// 启动调度器
func (s *TaskScheduler) Start() {
	// 你的代码
}

// 停止调度器
func (s *TaskScheduler) Stop() {
	// 你的代码
}
