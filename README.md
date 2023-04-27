# go-learning
初级入门，教程始于2023年4月25日，go版本1.20.3，开发工具GoLand

## 编码规范
- 1.go程序编写在.go为后缀的文件中
- 2.包名一般使用文件所在文件夹的名称，包名应该简洁、清晰且全小写
- 3.main函数只能编写在main包中
- 4.每一条语句后面可以不用编写分号(推荐)
- 5.如果没有编写分号,，一行只能编写一条语句
- 6.函数的左括号必须和函数名在同一行
- 7.导入包但没有使用包编译会报错
- 8.定义局部变量但没有使用变量编译也会报错
- 9.定义函数但没有使用函数不会报错
- 10.给方法、变量添加说明,尽量使用单行注释

## Golang各类型字节数
```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("int 占用的空间：", unsafe.Sizeof(int(0)))
	fmt.Println("int64 占用的空间：", unsafe.Sizeof(int64(0)))
	fmt.Println("boolean 占用的空间：", unsafe.Sizeof(true))
}
```
### 布尔类型
- true: 1个字节，1
- false：1个字节，0

### 整型
有符号整型包括int8、int16、int32和int64，而无符号整型包括uint8、uint16、uint32和uint64
- int8/uint8：1个字节     相当于java的 char
- int16/uint16：2个字节   相当于java的 short
- int32/uint32：4个字节   相当于java的 int
- int64/uint64：8个字节   相当于java的 long

### 浮点型
浮点型分为float32和float64两种类型
- float32：4个字节  相当于java的 float
- float64：8个字节  相当于java的 double

## 一些命令
```shell
# 初始化生成go.mod文件 go mod init 项目名
go mod init go-learning

```

## 建议学习顺序
- 【变量】learning_var.go
- 【常量】learning_const.go
- 【条件表达式】learning_if.go
- 【循环】learning_for.go
- 【switch】learning_switch.go  
- 【枚举】learning_enum.go
- 【函数】learning_func.go
- 【字符串与数字互转】learning_strconv.go
- 【http】learning_http.go
- 【延迟调用】learning_defer.go
- 【数组、切片】learning_array.go
- 【map】learning_map.go
- 【结构体、对象】learning_struct.go
- 【指针】learning_point.go
- 【接口】learning_interface.go
- 【异常】learning_error.go
- 【字符串相关】learning_string.go
- 【正则】learning_regexp.go
