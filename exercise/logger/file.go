package logger

// 往文件里面写日志

type FileLogger struct {
	level       LogLevel
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件保存的文件名
	maxFileSize int64  // 日志文件大小
}

func newFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	LogLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return &FileLogger{
		level:       LogLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
}

func (f FileLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= f.level
}

func (f FileLogger) Debug(format string, a ...interface{}) {
	if f.enable(DEBUG) {
		log(DEBUG, format, a...)
	}
}

func (f FileLogger) Trace(format string, a ...interface{}) {
	if f.enable(TRACE) {
		log(TRACE, format, a...)
	}
}

func (f FileLogger) Info(format string, a ...interface{}) {
	if f.enable(INFO) {
		log(INFO, format, a...)
	}
}

func (f FileLogger) Waring(format string, a ...interface{}) {
	if f.enable(WARING) {
		log(WARING, format, a...)
	}
}

func (f FileLogger) Error(format string, a ...interface{}) {
	if f.enable(ERROR) {
		log(ERROR, format, a...)
	}
}

func (f FileLogger) Fatal(format string, a ...interface{}) {
	if f.enable(FATAl) {
		log(FATAl, format, a...)
	}
}
