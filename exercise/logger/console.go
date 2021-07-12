package logger

import (
	"fmt"
	"time"
)

// 往终端写入日志

// Logger 日志对象/结构体
type consoleLogger struct {
	level LogLevel
}

// NewLog 构造函数
func NewLog(levelStr string) consoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return consoleLogger{
		level: level,
	}
}

func (c consoleLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= c.level
}

func (c consoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, fileName, lineNO := getCodeLine(3)
	lvStr, err := getLogString(lv)
	if err != nil {
		fmt.Printf("getLogString failed,err: %s\n", err)
		return
	}
	fmt.Printf("[%s] [%v] [%s %s %d] %s\n", now.Format("2006-01-02 15:04:05"), lvStr, fileName, funcName, lineNO, msg)
}

func (c consoleLogger) Debug(format string, a ...interface{}) {
	if c.enable(DEBUG) {
		c.log(DEBUG, format, a...)
	}
}

func (c consoleLogger) Trace(format string, a ...interface{}) {
	if c.enable(TRACE) {
		c.log(TRACE, format, a...)
	}
}

func (c consoleLogger) Info(format string, a ...interface{}) {
	if c.enable(INFO) {
		c.log(INFO, format, a...)
	}
}

func (c consoleLogger) Waring(format string, a ...interface{}) {
	if c.enable(WARING) {
		c.log(WARING, format, a...)
	}
}

func (c consoleLogger) Error(format string, a ...interface{}) {
	if c.enable(ERROR) {
		c.log(ERROR, format, a...)
	}
}

func (c consoleLogger) Fatal(format string, a ...interface{}) {
	if c.enable(FATAl) {
		c.log(FATAl, format, a...)
	}
}
