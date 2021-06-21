package main

import (
	"fmt"
	"os"
)

// os.OpenFilr() 函数能够以指定模式打开文件,从而实现文件写入相关功能
/*
func OpenFile (name string,flag int, perm FileMode) (*File, error) {
	....
}

name：要打开的文件名 flag：打开文件的模式。 模式有以下几种：

模式			含义
os.O_WRONLY	只写
os.O_CREATE	创建文件
os.O_RDONLY	只读
os.O_RDWR	读写
os.O_TRUNC	清空
os.O_APPEND	追加
perm：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。

*/
func main() {
	fileObj, err := os.OpenFile("./log.text", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("ioutil ReadFile file failed, err:%v\n", err)
		return
	}
	// write
	_, _ = fileObj.Write([]byte("法外狂徒张三"))
	// writeString
	_, _ = fileObj.WriteString("逃跑了\n")
	_ = fileObj.Close()
}
