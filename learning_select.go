package main

import (
	"fmt"
	"math/rand"
	"time"
)

// https://www.jianshu.com/p/de4bc02e7c72
// select 和 switch 类似，但是select语句只能用于信道的读写操作
func main() {
	demoSelect1()

	// 有时候我们会让main函数阻塞不退出，如http服务，我们会使用空的select{}来阻塞main goroutine
	// 这样主函数就永远阻塞住了，这里要注意上面一定要有一直活动的goroutine否则会报deadlock。

	// 两种死循环比较
	// for{} 使用100%CPU,因为它连续执行循环迭代.
	// select{} 使用接近0%的CPU因为它会导致go程到块,这意味着调度器把它睡觉,并且永远不会被唤醒.
}

func demoSelect1() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		a := rand.Intn(15)
		fmt.Printf("ch1休眠秒数%d\n", a)
		time.Sleep(time.Second * time.Duration(a))
		ch1 <- 1
	}()

	go func() {
		a := rand.Intn(15)
		fmt.Printf("ch2休眠秒数%d\n", a)
		time.Sleep(time.Second * time.Duration(a))
		ch2 <- 2
	}()

	select {
	case i := <-ch1:
		fmt.Printf("从ch1读取了数据%d\n", i)
	case j := <-ch2:
		fmt.Printf("从ch2读取了数据%d\n", j)
	case k := <-time.After(time.Second * 5):
		fmt.Printf("超过了5秒，两个channel都没有数据%v\n", k)
	}
}
