package main

import (
	"fmt"
	"net"
)

// tcp client
func main() {
	//1、与serve端简历连接
	dial, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		fmt.Println("dial 127.0.0.1:6666 failed", err)
		return
	}
	//	2、发送数据
	dial.Write([]byte("你好哇 ，李银河"))
	dial.Close()
}
