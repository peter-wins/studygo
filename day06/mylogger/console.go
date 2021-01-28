package mylogger

import (
	"fmt"
	"time"
)

//往终端写日志相关内容

// Logger 日志结构体
type Logger struct {
	Level LogLevel
}

// NewLog 构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func (l Logger) enable(logLevel LogLevel) bool {
	return logLevel >= l.Level
}
func log(lv LogLevel, msg string) {
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] - [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
}

// Debug 构造Debug日志级别的方法
func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		log(DEBUG, msg)
	}
}

// Info 构造Info日志级别的方法
func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		log(INFO, msg)
	}
}

// Warning 构造Warning日志级别的方法
func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		log(WARNING, msg)
	}
}

// Error 构造Error日志级别的方法
func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		log(ERROR, msg)
	}
}

// Fatal 构造Fatal日志级别的方法
func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		log(FATAL, msg)
	}
}
