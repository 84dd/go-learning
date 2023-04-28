package main

import "fmt"

// 管道是现成安全的，也是自带锁定功能的
// 声明：var 变量名 chan 数据类型   var ch1 chan int
// 初始化：变量名 = make(chan 数据类型, 容量)  ch1 = make(chan int, 3)
// 临时：ch1 := make(chan int)  可以不指定容量
// 注意：channel、map和切片，必须make后才能使用

// 定义管道
var gloCha = make(chan int, 3)
var gloExit = make(chan bool)

func main() {
	mainChannel()
	go goChannel()
	for {
		if num, ok := <-gloCha; ok {
			fmt.Println("遍历读到数据", num)
		} else {
			break
		}
	}
	<-gloExit

	// 单向管道和双向管道
	// 默认情况下所有管道都是双向了(可读可写)
	// 但是在企业开发中,我们经常需要用到将一个管道作为参数传递。
	// 在传递的过程中希望对方只能单向使用,要么只能写,要么只能读

	// 双向管道
	rwCh := make(chan int, 0)
	fmt.Println(len(rwCh))
	// 单向管道
	rCh := make(<-chan int, 0) // 只读
	wCh := make(chan<- int, 0) // 只写
	fmt.Println(len(rCh))
	fmt.Println(len(wCh))
	// 注意点:
	// 双向管道可以自动转换为任意一种单向管道
	// 单向管道不能转换为双向管道
}

// 主线程中的管道
func mainChannel() {
	// 定义管道
	ch1 := make(chan int, 3)
	// 往管道中写数据
	ch1 <- 111
	ch1 <- 222
	ch1 <- 333
	// 如果管道已满，再写入数据会报错
	//ch1 <- 444

	// 管道的本质是队列，先进先出
	num1 := <-ch1
	fmt.Println("读取到数据：", num1) // 111
	num2 := <-ch1
	fmt.Println("读取到数据：", num2) // 222
	num3 := <-ch1
	fmt.Println("读取到数据：", num3) // 333
	// 如果管道是空的，再读取会报错
	//num4 := <-ch1
	//fmt.Println("读取到数据：", num4)

	// 读取后，队列就空间了，可以再次写入
	ch1 <- 666
	ch1 <- 777
	ch1 <- 888

	// 关闭管道
	// 关闭管道后，只能读不能写
	// 如果在遍历前不关闭管道，会报错fatal error: all goroutines are asleep - deadlock!
	// 因为读到最后管道为空了，读空管道会报错
	close(ch1)

	// 管道遍历
	for {
		if num, ok := <-ch1; ok {
			fmt.Println("遍历读到数据", num)
		} else {
			break
		}
	}
	fmt.Println("数据读取完毕")
}

// 协程中的管道
func goChannel() {
	// 往管道中写数据
	gloCha <- 111
	gloCha <- 222
	gloCha <- 333
	// 在go程中的channel，当管道慢了，再写数据，并不会报错，而是阻塞
	gloCha <- 444
	gloCha <- 555
	gloCha <- 666
	gloCha <- 777
	close(gloCha)
	gloExit <- true
}
