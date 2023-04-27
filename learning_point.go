package main

import "fmt"

// point 指针
// 无论什么类型的指针，占用内存是固定的（32位机器4字节，64位机器8字节）
// 取地址：&
// 访问地址的值：*
func main() {
	// 普通指针
	num := 666
	pp := &num
	fmt.Println(pp)  // 打印地址 0xc0000180c8
	fmt.Println(*pp) // 打印地址的值 666

	// 数组指针
	arr := [3]int{1, 3, 5}
	pp2 := &arr  // 将arr的地址赋给指针
	pp2[1] = 777 // 指针地址改变为666，则指针对用的数组也同时修改了
	fmt.Println(arr[1])
	fmt.Println(pp2[1])

	// map指针
	map1 := map[string]int{"小明": 1, "小红": 2}
	mapPoint := &map1
	(*mapPoint)["小明"] = 2 // 对地址改值
	fmt.Println(map1)
	fmt.Println(*mapPoint)
	fmt.Println((*mapPoint)["小明"])

	// 结构体指针
	obj := struct {
		name string
		age  int
	}{
		name: "小明",
		age:  18,
	}
	objPoint := &obj
	(*objPoint).age = 18 // 对地址改值
	fmt.Println(obj)

	// 指针也可以作为参数或者返回值
	type Item struct {
		name string
		age  int
	}
	item := Item{name: "小明", age: 18}
	// 值传递
	changeAge1 := func(item Item) {
		item.age = 19
	}
	changeAge1(item)
	fmt.Println("调用 changeAge1 后，item.age = ", item.age) // 18 值传递，所以不会改变原来的值
	// 地址传递
	changeAge2 := func(item *Item) {
		item.age = 19
	}
	changeAge2(&item)
	fmt.Println("调用 changeAge2 后，item.age = ", item.age) // 19 地址传递
}
