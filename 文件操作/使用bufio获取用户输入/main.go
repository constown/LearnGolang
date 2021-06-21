package main

import (
	"bufio"
	"fmt"
	"os"
)

func userScan() {
	fmt.Print("content:")
	var s string
	fmt.Scanln(&s)
	fmt.Println("输入的内容是:", s)
}

func userBufIo() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("content:")
	s, _ = reader.ReadString('\n')
	fmt.Println("输入的内容是:", s)
}
func main() {
	userBufIo()
}
