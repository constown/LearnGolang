package main

import (
	"fmt"
	"net"
)

// tcp server端
func processConn(conn net.Conn) {
	var tmp [128]byte
	for {
		read, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("Read failed", err)
			return
		}
		fmt.Println(string(tmp[:read]))
	}
}

func main() {
	//1、本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:6666")
	if err != nil {
		fmt.Println("net listen on 127.0.0.1:6666 failed", err)
		return
	}
	//2、等待建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed", err)
			return
		}
		//	3、通信
		go processConn(conn)
	}
}
