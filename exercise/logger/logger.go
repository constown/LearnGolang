package logger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

// LogLevel 定义日志级别类型
type LogLevel uint16

// 定义日志的级别
const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARING
	ERROR
	FATAl
)

// 解析日志级别
func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToUpper(s)
	switch s {
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "INFO":
		return INFO, nil
	case "WARING":
		return WARING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAl":
		return FATAl, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

func getLogString(lv LogLevel) (string, error) {
	switch lv {
	case DEBUG:
		return "DEBUG", nil
	case TRACE:
		return "TRACE", nil
	case INFO:
		return "INFO", nil
	case WARING:
		return "WARING", nil
	case ERROR:
		return "ERROR", nil
	case FATAl:
		return "FATAl", nil
	default:
		err := errors.New("无效的日志级别")
		return "UNKNOWN", err
	}
}

func getCodeLine(n int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(n) //  这个n表示查找层次 ，当前层次为0
	if !ok {
		fmt.Printf("runtime.Caller failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1] // 分割一下不要包名
	fileName = path.Base(file)
	return
}
