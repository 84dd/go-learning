package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "http://wx.qlogo.cn/Vaz7vE1/64"
	resp1, err1 := http.Get(url) // 返回两个值，用两个变量分别接收
	fmt.Println(resp1, err1)
	_, err2 := http.Get(url) // 返回两个值，但第一个值用 _ 下划线，表示丢弃不需要的值
	fmt.Println(err2)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintf(writer, "asdaaass")
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe(":8899", nil)
	if err != nil {
		return
	}
}
