package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		return
	}
	var data [1024]byte
	for {
		n, addr, err2 := conn.ReadFromUDP(data[:])
		if err2 != nil {
			return
		}
		fmt.Println(data[:n])
		reply := strings.ToUpper(string(data[:n]))
		conn.WriteToUDP([]byte(reply), addr)
	}
}
