package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 打开文件读内容
// os.Open() 函数能打开一个文件,返回一个 *file 和一个 err, 对得到的文件实例,使用 close() 方法可以关闭文件
// 为了防止文件忘记关闭,我们通常使用defer语句注册文件关闭语句

func readFile() {
	fileObj, err := os.Open("./log.text")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {
			fmt.Printf("close file failed, err:%v\n", err)
			return
		}
	}(fileObj)
	var tmp = make([]byte, 128) // 指定读的长度
	for {
		n, err := fileObj.Read(tmp[:])
		if n == 0 {
			return
		}
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("Read file failed, err:%v\n", err)
			return
		}
		fmt.Println(22)
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
	}
}

// 利用Bufio读文件
func readFileByBufio() {
	fileObj, err := os.Open("./log.text")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {
			fmt.Printf("close file failed, err:%v\n", err)
			return
		}
	}(fileObj)
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("ReadString file failed, err:%v\n", err)
			return
		}
		fmt.Print(line)
	}
}

// 读整个文件,自动打开关闭文件
func readFileByIoutil() {
	ret, err := ioutil.ReadFile("./log.text")
	if err != nil {
		fmt.Printf("ioutil ReadFile file failed, err:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	//readFile()
	//readFileByBufio()
	readFileByIoutil()
}
