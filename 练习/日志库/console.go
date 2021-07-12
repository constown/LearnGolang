package mylogger

import "fmt"

// 往终端写入日志

// Logger 日志对象
type Logger struct {
}

// NewLog 构造函数
func NewLog() Logger {
	return Logger{}
}

func (l Logger) Debug(msg string) {
	fmt.Println(msg)
}

func (l Logger) Info(msg string) {
	fmt.Println(msg)
}

func (l Logger) Waring(msg string) {
	fmt.Println(msg)
}

func (l Logger) Error(msg string) {
	fmt.Println(msg)
}

func (l Logger) Fatal(msg string) {
	fmt.Println(msg)
}
