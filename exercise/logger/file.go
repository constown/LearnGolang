package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日志

type FileLogger struct {
	level       LogLevel
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64 // 日志文件大小
}

// NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	LogLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		level:       LogLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err1 := f1.initFile() // 按照文件路径和文件名打开
	if err1 != nil {
		panic(err1)
	}
	return f1
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err: %v\n", err)
		return err
	}
	errFileObj, err1 := os.OpenFile(fullFileName+".err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		fmt.Printf("open log file failed, err: %v\n", err1)
		return err1
	}
	// 日志文件都打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	// 在这里判断日志的级别
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNO := getCodeLine(3)
		lvStr, _ := getLogString(lv)
		fmt.Fprintf(f.fileObj, "[%s] [%v] [%s %s %d] %s\n", now.Format("2006-01-02 15:04:05"), lvStr, fileName, funcName, lineNO, msg)
		// 如果日志级别大于 ERROR 还需要在 err 日志文件中记录一遍
		if lv >= ERROR {
			fmt.Fprintf(f.errFileObj, "[%s] [%v] [%s %s %d] %s\n", now.Format("2006-01-02 15:04:05"), lvStr, fileName, funcName, lineNO, msg)
		}
	}
}

func (f FileLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= f.level
}

func (f FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

func (f FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, format, a...)
}

func (f FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

func (f FileLogger) Waring(format string, a ...interface{}) {
	f.log(WARING, format, a...)
}

func (f FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

func (f FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAl, format, a...)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
