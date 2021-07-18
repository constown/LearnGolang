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

// NewFileLogger 构造函数日志文件结构体
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

// 文件初始化
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

// 检查文件大小，如果超过了就表示应该切割日志了
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return false
	}
	return fileInfo.Size() >= f.maxFileSize
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	nowStr := time.Now().Format("20060102250405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())           // 拿到当前日志文件的完整路径
	newLogName := fmt.Sprintf("%s.back%s.log", logName, nowStr) // 拼接一个备份的日志文件的名字
	// 1、关闭当前日志文件
	file.Close()
	// 2、rename备份一下
	os.Rename(logName, newLogName)
	// 3、打开一个新的日志文件
	fileObj, err1 := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err1 != nil {
		fmt.Printf("open new log file failed,err:%v\n", err)
		return nil, err1
	}
	// 4、将打开的新的日志文件对象赋值给 f.fileObj
	return fileObj, nil
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	// 在这里判断日志的级别
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNO := getCodeLine(3)
		lvStr, _ := getLogString(lv)
		if f.checkSize(f.fileObj) {
			// 如果为真表示需要切割
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}

		fmt.Fprintf(f.fileObj, "[%s] [%v] [%s %s %d] %s\n", now.Format("2006-01-02 15:04:05"), lvStr, fileName, funcName, lineNO, msg)
		// 如果日志级别大于 ERROR 还需要在 err 日志文件中记录一遍
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				// 如果为真表示需要切割
				newFile, err := f.splitFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
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
