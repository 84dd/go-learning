package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "120.79.51.157:9090")
	if err != nil {
		fmt.Printf("链接TCP失败, err : %v\n", err.Error())
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("获取键盘输入失败, err: %v\n", err)
			break
		}
		trimmedInput := strings.TrimSpace(input)
		// 发送数据
		// TCP 包末尾加 \n 特殊字符用于服务端拆包
		if _, err = conn.Write([]byte(trimmedInput + "\n")); err != nil {
			fmt.Printf("发送数据失败 , err : %v\n", err)
			break
		}
		// 接受数据
		var recData = make([]byte, 1024)
		if _, err = conn.Read(recData); err != nil {
			fmt.Printf("接收数据失败 , err : %v\n", err)
			break
		}
		fmt.Printf("接收到数据 : %v\n", string(recData))
	}
}
