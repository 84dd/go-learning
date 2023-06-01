package main

import (
	"errors"
	"fmt"
)

func main() {
	// 创建异常的方式有两种
	// 方式一：使用errors.New 函数
	err1 := errors.New("这里报错了")
	fmt.Println(err1)

	// 方式二：利用 fmt.Error 其底层实现原理就是在内部调用了 errors.New
	err2 := fmt.Errorf("这里还是报错了，行数是%v", 99)
	fmt.Println(err2)

	// error：相当于java中抛出了异常，并catch住了
	num1, err3 := divError(10, 0)
	fmt.Println(num1)
	fmt.Println(err3)

	{
		//defer func() {
		//	if err := recover(); err != nil {
		//		// 捕获了异常，程序恢复
		//		fmt.Println(err)
		//	}
		//}()
		// panic：程序会终止，相当于java中仅仅抛出了异常
		// TODO 建议开放下面注释试一下效果
		//num2 := divPanic(10, 0)
		//fmt.Println(num2)
	}

	// 触发panic有两种方式
	// 方式一：手动panic    panic("报错了")
	// 方式二：自动panic    fmt.Println(10/0)
	// fmt.Println(10 / 0)

	num3 := divPanicRecover(10, 0)
	fmt.Println(num3)

	defer func() {
		if err := recover(); err != nil {
			// 捕获了异常，程序恢复
			fmt.Println(err)
		}
	}()
	// 多个异常只会捕获第一个
	// 如果想演示 divDeferPanic，需要把 divMultiPanic 注释
	num4 := divMultiPanic(10, 0)
	fmt.Println(num4)

	// 如果想演示 divDeferPanic，需要把 divMultiPanic 注释
	num5 := divDeferPanic(10, 0)
	fmt.Println(num5)

	// 如果演示了 divMultiPanic 或 divDeferPanic，则下面的话不能输出
	// 原因：defer 捕获的异常，是在本方法结束后才捕获的，而下面的输出是正常结束才会输出
	// 如果 panic，并且用 defer 捕获，说明运行遇上 panic ，本方法立即退出，并在退出时捕获panic
	// 相当于java中的 try catch finally
	fmt.Println("能看到这句话，证明程序正常结束")
}

// 简单输出异常
func divError(a int, b int) (res float32, err error) {
	if b == 0 {
		//err = errors.New("这里报错了")
		err = fmt.Errorf("除数不能为0")
	} else {
		res = float32(a / b)
	}
	return
}

// 抛出异常，但不捕获
func divPanic(a int, b int) (res float32) {
	if b == 0 {
		// 程序终止
		panic("除数不能为0")
	} else {
		res = float32(a / b)
	}
	return
}

// 抛出异常，并捕获异常
func divPanicRecover(a int, b int) (res float32) {
	// 利用defer特性，捕获异常
	// 注意：一定要在 panic 前，定义 defer
	defer func() {
		// 利用 recover() 获取异常，当异常不为空，进入判断
		if err := recover(); err != nil {
			// 捕获了异常，程序恢复
			res = -1
			fmt.Println(err)
		}
	}()
	if b == 0 {
		// 程序终止
		panic("除数不能为0")
	} else {
		res = float32(a / b)
	}
	return
}

// 多 panic 的情况
func divMultiPanic(a int, b int) (res float32) {
	if b == 0 {
		// 程序终止，并且故意多个 panic
		panic("111除数不能为0") // 只有这个会被捕获
		panic("222除数不能为0")
		panic("333除数不能为0")
		panic("444除数不能为0")
	} else {
		res = float32(a / b)
	}
	return
}

// defer panic
func divDeferPanic(a int, b int) (res float32) {
	if b == 0 {
		defer func() {
			// 在defer中抛出的异常，程序也会终止，而且 defer 中的panic，只能在defer中被捕获
			panic("defer除数不能为0")
		}()
	} else {
		res = float32(a / b)
	}
	return
}
