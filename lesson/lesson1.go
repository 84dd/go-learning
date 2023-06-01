package lesson

// 1.在你的电脑上安装 golang
// 2.在你的电脑上安装 vscode 或者 goland
// 3.IDE 配置好 golang 开发环境
// 4.在当前项目的 cmd 目录下创建以你名字拼音命名的目录，再在里面新增 main.go 文件
// 5.在 main.go 文件内创建 main 方法，打印出 hello world
// 思考题：
// 1.一定要用 main.go 这个文件名吗？ 用其他名字是否可行
// 2.文件内的 package main 可以把 main 改成其他的名字后，然后正常编译执行这个文件吗？
// 3.函数名为什么一定用 main 才可以运行

// 参考答案
// 思考题：
// 1.一定要用 main.go 这个文件名吗？ 用其他名字是否可行
// 可以用其他文件名，比如 xxx.go

// 2.文件内的 package main 可以把 main 改成其他的名字后，然后正常编译执行这个文件吗？
// 可以把main包换成其他名字，可以正常编译。但这样就丧失了运行的权利。

// 3.函数名为什么一定用 main 才可以运行
// go定义了 package main 的 main 方法是程序的入口
// 如果其他包下的方法也想运行，可以参考 Test 的方法，即 文件名为  xxx_test.go  方法名为  Testxxxx
