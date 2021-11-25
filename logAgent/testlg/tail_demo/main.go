package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

// tailf的用法
func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的那个位置开始读
		MustExist: false,                                // 是否必须存在
		Poll:      true,                                 //
	}
	tails, err := tail.TailFile(fileName, config) // 打开文件
	if err != nil {
		fmt.Println("tail file fauked err:", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-tails.Lines // 一行一行的读日志
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg:", line.Text)
	}
}
