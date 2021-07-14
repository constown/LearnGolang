package main

import (
	logger "LearnGolang/exercise/logger"
)

// 测试日志库
func main() {
	log := logger.NewFileLogger("debug", "./", "zhoulinewan", 10*1024*1024)
	log.Debug("这是一条Debug日志")
	log.Info("这是一条info日志")
	log.Waring("这是一条Waring日志")
	log.Error("这是一条Error日志")
	log.Fatal("这是一条Fatal日志")
}
