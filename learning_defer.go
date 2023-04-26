package main

import "fmt"

func main() {
	// defer 延迟调用，相当于java中的 finally
	// 注意：一个方法中声明了多个defer, 那么 defer 是按顺序从后往前执行的

	// https://blog.csdn.net/maihilton/article/details/108887304

	fmt.Println("defer1 => ", defer1())
	fmt.Println("defer2 => ", defer2())
	fmt.Println("defer3 => ", defer3())
}

// 不是 10 , 是 11
func defer1() (res int) {
	defer func() {
		// 2. res = 10 + 1 = 11
		res++
	}()
	// 1. res = 10
	// 3. return res
	return 10
}

// 不是15, 是10
func defer2() (res int) {
	// 1. sb = 10
	sb := 10
	defer func() {
		// 3. sb = 15, 但是 res = 10
		sb += 5
	}()
	// 2. res = sb = 10
	// 4. return res(10)
	return sb
}

// 不是12，是10
func defer3() (res int) {
	// 1. res=2
	res = 2
	defer func(res int) {
		// 3. 内部res为形参，不影响外边的值 res=2+2=4
		res += 2
		fmt.Println("内部 res ", res) // 4
	}(res) // defer 参数的值是在声明的时候确定的，也就是只有 defer 之前的语句会影响这个值
	// 2. res = 10
	// 4. return res(10)
	return 10
}
