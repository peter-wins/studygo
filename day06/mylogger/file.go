package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里写日志
type FileLogger struct {
	Level       LogLevel
	filePath    string //日志保存的路径
	fileName    string //日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

// NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() //按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

// 根据指定的日志文件路径和文件名打开日志
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed! err:%v", err)
		return err
	}

	fullFileName = path.Join(f.filePath, f.fileName)
	errfileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed! err:%v", err)
		return err
	}
	// 日志文件都已经打开了，开始记录日志
	f.fileObj = fileObj
	f.errFileObj = errfileObj
	return nil
}

// 判断是否记录该日志
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

// 判断文件是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Get file info failed, err%v\n", err)
		return false
	}
	//如果当前文件大小 大于等于 日志文件的最大值 就应该返回true
	return fileInfo.Size() >= f.maxFileSize
}

// 切割文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	//需要切割日志文件
	newStr := time.Now().Format("20060102030405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Get file info failed, err%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())      //拿到当前的日志文件完整的路径
	NewLogName := fmt.Sprintf("%s.bak%s", logName, newStr) //拼接一个日志文件备份的名字
	//1.关闭当前的日志文件
	file.Close()
	//2.备份一下 rename
	os.Rename(logName, NewLogName)
	//3.打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err%v\n", err)
		return nil, err
	}
	//4.将打开的新日志文件对象赋值给 f.fileObj
	return fileObj, nil
}

// 记录日志的方法
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		now := time.Now()
		msg := fmt.Sprintf(format, a...)
		funcName, fileName, lineNo := getInfo(3)
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] - [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)

		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				newFile, err := f.splitFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
			//如果要记录的日志大于等于ERROR级别，我还要在err中记录一下
			fmt.Fprintf(f.errFileObj, "[%s] - [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		}
	}
}

// Debug 构造Debug日志级别的方法
func (f *FileLogger) Debug(format string, a ...interface{}) {

	f.log(DEBUG, format, a...)
}

// Info 构造Info日志级别的方法
func (f *FileLogger) Info(format string, a ...interface{}) {

	f.log(INFO, format, a...)

}

// Warning 构造Warning日志级别的方法
func (f *FileLogger) Warning(format string, a ...interface{}) {

	f.log(WARNING, format, a...)

}

// Error 构造Error日志级别的方法
func (f *FileLogger) Error(format string, a ...interface{}) {

	f.log(ERROR, format, a...)

}

// Fatal 构造Fatal日志级别的方法
func (f *FileLogger) Fatal(format string, a ...interface{}) {

	f.log(FATAL, format, a...)

}

// Close 关闭文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
