package lesson

import (
	"os"
	"strings"
	"sync"
)

// 参考 amos 分享的《Go 语言圣经》第八章《Goroutines和Channels》的内容
// 在你个人目录下创建 lesson14.go ，然后完成以下练习题：
// 现在需要你通过使用 goroutine、channel、sync.WaitGroup 来实现一个函数 FileNameSearch
// 该函数可以使用多 goroutine 递归遍历（目录里面的目录也要遍历）你指定目录 d 下的所有文件名，如果文件名包含指定的关键字 k ，则将其放入通道 r，
// 最后将所有符合条件的文件名字（包含文件的路径）放在一个字符串切片中，并返回
// 比如
// 第 14 课
// dirPath := "/Users/amos/Downloads" // 搜索的目录路径
// keyword := ".jar"                  // 搜索的关键字
// fmt.Println(lesson.FileNameSearch(dirPath, keyword))

// 结果打印 [/Users/amos/Downloads/dubbo-consumer-1.0-SNAPSHOT-jar-with-dependencies.jar /Users/amos/Downloads/dubbo-consumer-1.0-SNAPSHOT.jar]

// searchFiles 并发搜索文件 ----- 递归函数
// - d: directory 文件目录名
// - k: keyword 需要搜索的关键字
// - r: result 搜索出来的匹配结果并放入通道
// - wg: goroutine 计数器 它可以等待其他 goroutine 结束
func searchFiles(d string, k string, r chan<- string, wg *sync.WaitGroup) {
	// 将路径的最后一个文件分隔符去掉，比如 d:\\file\\ -> d:\\file
	d = strings.TrimSuffix(d, string(os.PathSeparator))
	// 离开方法后，计数器减一
	defer wg.Done()
	fs, _ := os.ReadDir(d)
	for i := 0; i < len(fs); i++ {
		f := fs[i]
		path := d + string(os.PathSeparator) + f.Name()
		if f.IsDir() {
			// 如果是文件夹，计数器加1，再递归
			wg.Add(1)
			go searchFiles(path, k, r, wg)
		} else {
			// 如果是文件，则进入文件名匹配阶段
			searchFileName(path, k, r)
		}
	}
}

// searchFileName 搜索文件名字  ----- 符合条件的文件名则需要放到通道里
// - f 文件名字
// - k keyword 搜索关键子
// - r result 结果传输通道
func searchFileName(f string, k string, r chan<- string) {
	// 通过文件分隔符，将文件路径分隔成数组
	arr := strings.Split(f, string(os.PathSeparator))
	// 取文件名
	fn := arr[len(arr)-1]
	if len(k) == 0 || strings.Index(fn, k) > 0 {
		// 匹配规则：关键字为空 或 文件名包含关键字
		r <- f
	}
}

// FileNameSearch 文件名搜索函数  ------ 真正暴露给外部调用的函数
// - d directory 目录名
// - k keyword 搜索关键字
// - return : file string slice 搜索到的匹配文件字符串切片
func FileNameSearch(d string, k string) []string {
	arr := [...]string{}
	slice := arr[:]

	// 判断
	s, err := os.Stat(d)
	if err != nil || !s.IsDir() {
		// 如果目录路径不存在，或传参目录为文件，则不用搜索了，直接返回空
		return slice
	}

	// 创建无缓冲区的channel
	r := make(chan string)
	// 创建等待组
	wg := sync.WaitGroup{}
	// 计数器加1
	wg.Add(1)
	go searchFiles(d, k, r, &wg)

	// 开启协程读取数据，否则会出现：fatal error: all goroutines are asleep - deadlock!
	go func(r chan string) {
		for v := range r {
			// 不断读取内容放进切片中
			slice = append(slice, v)
		}
	}(r)

	// 同步等待
	wg.Wait()
	return slice
}
